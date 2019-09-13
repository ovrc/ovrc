import React, { useEffect, useState } from "react";
import { render } from "react-dom";
import { Router } from "@reach/router";
import UserContext from "./UserContext";

import PrivateRoute from "./PrivateRoute";
import Loading from "./Loading";
import { api_request } from "./Helpers";

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

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {user === null ? (
        <Loading />
      ) : (
        <Router>
          <PrivateRoute as={Dashboard} path="/" />
          <Test path="/test" />
        </Router>
      )}
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
