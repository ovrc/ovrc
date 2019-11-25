import React, { useState, useEffect } from "react";
import { api_request } from "./Helpers";

const HTTPMonitoringList = () => {
  const [monitors, setMonitors] = useState([]);

  useEffect(() => {
    api_request("/monitoring/http", "GET").then(res => {
      if (res.status === "success") {
        setMonitors(res.data.monitors);
      }
    });
  }, []);

  if (monitors.length > 0) {
    return (
      <div>
        <div className="buttons is-pulled-right	">
          <button className="button is-primary">Add</button>
        </div>
        <table className="table">
          <thead>
            <tr>
              <th>Endpoint</th>
              <th>Avg. Response Time</th>
            </tr>
          </thead>
          <tbody>
            {monitors.map(function(monitor) {
              return (
                <tr>
                  <td>
                    <b>{monitor.method}</b> {monitor.endpoint}
                  </td>
                  <td>{monitor.avg_total_ms}ms</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    );
  }

  return "Loading...";
};

export default HTTPMonitoringList;
