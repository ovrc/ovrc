import React, { useState, useEffect } from "react";
import { api_request } from "./Helpers";

const HTTPMonitoringList = () => {
  const [monitors, setMonitors] = useState([]);
  const [loaded, setLoaded] = useState(null);

  useEffect(() => {
    api_request("/monitoring/http", "GET").then(res => {
      if (res.status === "success") {
        if (res.data.monitors && res.data.monitors.length > 0) {
          setMonitors(res.data.monitors);
          setLoaded(true);
        } else {
          // No results.
          setLoaded(false);
        }
      }
    });
  }, []);

  if (loaded === null) {
    return "Loading...";
  }

  if (loaded === false) {
    return "No HTTP monitors found!";
  }

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
              <th>Avg. Total Time</th>
              <th>Avg. Total Time</th>
              <th>Avg. Total Time</th>
              <th>Avg. Total Time</th>
              <th>Avg. Total Time</th>
            </tr>
          </thead>
          <tbody>
            {monitors.map(function(monitor) {
              return (
                <tr key={monitor.id}>
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
};

export default HTTPMonitoringList;
