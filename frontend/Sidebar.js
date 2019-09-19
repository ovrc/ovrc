import React from "react";

const Sidebar = () => {
  return (
    <aside className="menu">
      <p className="menu-label">General</p>
      <ul className="menu-list">
        <li>
          <a className="is-active">Dashboard</a>
        </li>
      </ul>
      <p className="menu-label">Administration</p>
      <ul className="menu-list">
        <li>
          <a>Users</a>
        </li>
      </ul>
    </aside>
  );
};

export default Sidebar;
