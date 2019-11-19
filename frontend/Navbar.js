import React from "react";
import { Link } from "@reach/router";
import { api_request } from "./Helpers";
import UserContext from "./UserContext";

const Navbar = () => {
  function logout() {
    api_request("/auth/logout", "GET").then(res => {
      if (res.status === "success") {
        window.location.replace("../");
      }
    });
  }

  return (
    <section className="hero is-dark">
      <div className="hero-head">
        <nav
          className="navbar is-dark"
          role="navigation"
          aria-label="main navigation"
        >
          <div className="navbar-brand">
            <Link className="navbar-item" to="/">
              ovrc
            </Link>

            <a
              role="button"
              className="navbar-burger burger"
              aria-label="menu"
              aria-expanded="false"
              data-target="navbarBasicExample"
            >
              <span aria-hidden="true"></span>
              <span aria-hidden="true"></span>
              <span aria-hidden="true"></span>
            </a>
          </div>

          <div className="navbar-menu">
            <div className="navbar-end">
              <div className="navbar-item">
                Logged in as &nbsp;
                <UserContext.Consumer>
                  {context => <strong>{context.user.username}</strong>}
                </UserContext.Consumer>
                &nbsp;
                <div className="buttons">
                  <i className="button is-primary" onClick={logout}>
                    <strong>Logout</strong>
                  </i>
                </div>
              </div>
            </div>
          </div>
        </nav>
      </div>
    </section>
  );
};

export default Navbar;
