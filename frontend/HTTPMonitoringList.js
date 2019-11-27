import React, { useEffect, useState } from "react";
import { api_request } from "./Helpers";

const HTTPMonitoringList = () => {
  const [monitors, setMonitors] = useState([]);
  const [loaded, setLoaded] = useState(null);
  const [showAddModal, setShowAddModal] = useState(false);
  const [disableAddButton, setDisableAddButton] = useState(false);
  const [method, updateMethod] = useState("");
  const [url, updateUrl] = useState("");

  useEffect(() => {
    // Grab initial data on load.
    updateMonitoringList();

    // Update data every 15 seconds.
    setInterval(function() {
      updateMonitoringList();
    }, 15000);
  }, []);

  function updateMonitoringList() {
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
  }

  // Shows or hides the modal, depending on what it is currently set to.
  function handleShowAddModal() {
    setShowAddModal(!showAddModal);
  }

  function addMonitor() {
    setDisableAddButton(true);
    api_request("/monitoring/http", "POST", {
      method: method,
      url: url
    }).then(res => {
      setDisableAddButton(false);
      setShowAddModal(false);
      updateUrl("");
      updateMethod("");
    });
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
              <p className="modal-card-title">Add New Monitor</p>
              <button
                className="delete"
                aria-label="close"
                onClick={handleShowAddModal}
              ></button>
            </header>
            <form
              onSubmit={e => {
                e.preventDefault();
                addMonitor();
              }}
            >
              <section className="modal-card-body">
                <div className="field is-horizontal">
                  <div className="field-label is-normal">
                    <label className="label">Method</label>
                  </div>
                  <div className="field-body">
                    <div className="field">
                      <p className="control">
                        <div className="select">
                          <select
                            onChange={e => updateMethod(e.target.value)}
                            required
                          >
                            <option disabled selected value=""></option>
                            <option value="GET">GET</option>
                          </select>
                        </div>
                      </p>
                    </div>
                  </div>
                </div>
                <div className="field is-horizontal">
                  <div className="field-label is-normal">
                    <label className="label">URL</label>
                  </div>
                  <div className="field-body">
                    <div className="field">
                      <p className="control">
                        <input
                          className="input"
                          type="text"
                          placeholder="https://www.google.com/"
                          onChange={e => updateUrl(e.target.value)}
                          required
                          value={url}
                        />
                      </p>
                    </div>
                  </div>
                </div>
              </section>
              <footer className="modal-card-foot">
                <button
                  disabled={disableAddButton}
                  className={
                    "button is-success " +
                    (disableAddButton ? "is-loading" : null)
                  }
                >
                  Add Monitor
                </button>
                <button className="button" onClick={handleShowAddModal}>
                  Cancel
                </button>
              </footer>
            </form>
          </div>
        </div>
      </div>
    );
  }
};

export default HTTPMonitoringList;
