import React from "react";

const HTTPMonitoringList = () => {
  return (
    <div>
      <div className="buttons is-pulled-right	">
        <button className="button is-primary">Add</button>
      </div>
      <table className="table">
        <thead>
          <tr>
            <th>Endpoint</th>
          </tr>
        </thead>
        <tbody></tbody>
      </table>
    </div>
  );
};

export default HTTPMonitoringList;
