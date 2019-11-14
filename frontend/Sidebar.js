import React from "react";
import { Link } from "@reach/router";

const Sidebar = () => {
  const isActive = ({ isCurrent }) => {
    return isCurrent ? { className: "is-active" } : null;
  };

  return (
    <aside className="menu">
      <p className="menu-label">General</p>
      <ul className="menu-list">
        <li>
          <Link getProps={isActive} to="/">
            Dashboard
          </Link>
        </li>
      </ul>
      <p className="menu-label">Administration</p>
      <ul className="menu-list">
        <li>
          <Link getProps={isActive} to="users">
            Users
          </Link>
        </li>
      </ul>
    </aside>
  );
};

export default Sidebar;
