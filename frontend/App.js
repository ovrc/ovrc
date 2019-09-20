import React, { useEffect, useState } from "react";
import { render } from "react-dom";
import { Router } from "@reach/router";
import UserContext from "./UserContext";

import PrivateRoute from "./PrivateRoute";
import Sidebar from "./Sidebar";
import Loading from "./Loading";
import { api_request } from "./Helpers";
import Login from "./Login";
import Navbar from "./Navbar";

const App = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Small time out so the loading page doesn't just flash the screen.
    const timer = setTimeout(() => {
      api_request("/users/me", "GET").then(res => {
        if (res.status === "success") {
          setUser(true);
        } else {
          setUser(false);
        }
      });
    }, 300);
    return () => clearTimeout(timer);
  }, []);

  // When the status of the user is yet to be checked.
  if (user === null) {
    return <Loading />;
  }

  // If user validation failed (not logged in/logged out).
  if (user === false) {
    return (
      <UserContext.Provider value={{ user, setUser }}>
        <Login />
      </UserContext.Provider>
    );
  }

  return (
    <UserContext.Provider value={{ user, setUser }}>
      <Navbar />
      <div className="columns">
        <div className="column is-2 aside">
          <Sidebar />
        </div>
        <div className="column aside">
          <div className="box">
            <Router>
              <PrivateRoute as={Dashboard} path="/" />
              <Test path="/test" />
            </Router>
          </div>
        </div>
      </div>
    </UserContext.Provider>
  );
};

const Dashboard = () => {
  return "Protected Dashboard!";
};

const Test = () => {
  return "test";
};

render(<App />, document.getElementById("root"));
