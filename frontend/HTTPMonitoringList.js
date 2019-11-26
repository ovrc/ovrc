import React, { useEffect, useState } from "react";
import { api_request } from "./Helpers";

const HTTPMonitoringList = () => {
  const [monitors, setMonitors] = useState([]);
  const [loaded, setLoaded] = useState(null);
  const [showAddModal, setShowAddModal] = useState(false);

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

  function handleShowAddModal() {
    setShowAddModal(true);
  }

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
          <button className="button is-primary" onClick={handleShowAddModal}>
            Add
          </button>
        </div>
        <table className="table">
          <thead>
            <tr>
              <th>Endpoint</th>
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
        <div
          id="add_monitor"
          className={"modal " + (showAddModal ? "is-active" : null)}
        >
          <div className="modal-background"></div>
          <div className="modal-card">
            <header className="modal-card-head">
              <p className="modal-card-title">Modal title</p>
              <button className="delete" aria-label="close"></button>
            </header>
            <section className="modal-card-body"></section>
            <footer className="modal-card-foot">
              <button className="button is-success">Save changes</button>
              <button className="button">Cancel</button>
            </footer>
          </div>
        </div>
      </div>
    );
  }
};

export default HTTPMonitoringList;
