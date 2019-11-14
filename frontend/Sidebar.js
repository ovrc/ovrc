import React from "react";
import { Link } from "@reach/router";

const Sidebar = () => {
  return (
    <aside className="menu">
      <p className="menu-label">General</p>
      <ul className="menu-list">
        <li>
          <Link className="is-active" to="/">
            Dashboard
          </Link>
        </li>
      </ul>
      <p className="menu-label">Administration</p>
      <ul className="menu-list">
        <li>
          <Link to="users">Users</Link>
        </li>
      </ul>
    </aside>
  );
};

export default Sidebar;
