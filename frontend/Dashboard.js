import React, { useState, useEffect } from "react";
import { api_request } from "./Helpers";

const Dashboard = () => {
  const [dashboardTiles, setDashboardTiles] = useState(null);

  useEffect(() => {
    api_request("/dashboard/tiles", "GET").then(res => {
      if (res.status === "success") {
        setDashboardTiles(res.data);
      }
    });
  }, []);

  if (!dashboardTiles) {
    return "Loading...";
  }

  return (
    <div className="tile is-ancestor">
      <div className="tile is-parent ">
        <article className="tile is-child box notification is-info ">
          <p className="title">{dashboardTiles.http_monitors}</p>
          <p className="subtitle">HTTP Monitors</p>
        </article>
      </div>
    </div>
  );
};

export default Dashboard;
