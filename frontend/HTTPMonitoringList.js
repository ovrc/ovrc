import React, { useEffect, useState } from "react";
import { api_request } from "./Helpers";
import { Sparklines, SparklinesLine, SparklinesSpots } from "react-sparklines";
import prettyMilliseconds from "pretty-ms";

const HTTPMonitoringList = () => {
  const [monitors, setMonitors] = useState([]);
  const [loaded, setLoaded] = useState(false);
  const [showAddModal, setShowAddModal] = useState(false);
  const [disableAddButton, setDisableAddButton] = useState(false);
  const [method, updateMethod] = useState("");
  const [url, updateUrl] = useState("");
  const [hideNotification, setHideNotification] = useState(true);
  const [periodSelectValue, setPeriodSelectValue] = useState("hour24");

  useEffect(() => {
    updateMonitoringList();
  }, [periodSelectValue]);

  function updateMonitoringList() {
    setLoaded(false);
    setMonitors([]);
    api_request("/monitoring/http", "GET", { period: periodSelectValue }).then(
      res => {
        if (res.status === "success") {
          setLoaded(true);
          if (res.data.monitors && res.data.monitors.length > 0) {
            setMonitors(res.data.monitors);
          }
        }
      }
    );
  }

  // Shows or hides the modal, depending on what it is currently set to.
  function handleShowAddModal() {
    setShowAddModal(!showAddModal);
  }

  function handleNotification() {
    setHideNotification(!hideNotification);
  }

  function addMonitor() {
    setDisableAddButton(true);
    api_request("/monitoring/http", "POST", {
      method: method,
      url: url
    }).then(res => {
      setDisableAddButton(false);
      setShowAddModal(false);
      setHideNotification(false);
      updateUrl("");
      updateMethod("");
    });
  }

  return (
    <div>
      <div
        className={
          "notification is-primary " + (hideNotification ? "is-hidden" : null)
        }
      >
        <button className="delete" onClick={handleNotification}></button>
        New monitor added! Please wait for data to refresh...
      </div>

      <div className="buttons is-pulled-right	">
        <button className="button is-primary" onClick={handleShowAddModal}>
          Add
        </button>
      </div>
      <div className="select">
        <select
          onChange={e => setPeriodSelectValue(e.target.value)}
          value={periodSelectValue}
        >
          <option value="hour1">Last 1h</option>
          <option value="hour3">Last 3h</option>
          <option value="hour6">Last 6h</option>
          <option value="hour12">Last 12h</option>
          <option value="hour24">Last 24h</option>
        </select>
      </div>
      <hr />
      {monitors.length > 0 ? (
        <table className="table">
          <thead>
            <tr>
              <th>Endpoint</th>
              <th>Avg. Total Time</th>
              <th>Min Total Time</th>
              <th>Max Total Time</th>
              <th>
                {/* TODO: Find an alternative to this? */}
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
              </th>
            </tr>
          </thead>
          <tbody>
            {monitors.map(function(monitor) {
              return (
                <tr key={monitor.id}>
                  <td>
                    <span className="tag is-primary is-light">
                      <b>{monitor.method}</b>
                    </span>{" "}
                    {monitor.endpoint}
                  </td>
                  <td>{prettyMilliseconds(monitor.avg_total_ms)}</td>
                  <td>{prettyMilliseconds(monitor.min_total_ms)}</td>
                  <td>{prettyMilliseconds(monitor.max_total_ms)}</td>
                  <td>
                    <Sparklines data={monitor.last_entries}>
                      <SparklinesLine color="green" />
                      <SparklinesSpots />
                    </Sparklines>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      ) : loaded === false ? (
        "Loading..."
      ) : (
        "No monitors found!"
      )}
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
                    <div className="control">
                      <div className="select">
                        <select
                          onChange={e => updateMethod(e.target.value)}
                          required
                          value=""
                        >
                          <option disabled value=""></option>
                          <option value="GET">GET</option>
                        </select>
                      </div>
                    </div>
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
};

export default HTTPMonitoringList;
