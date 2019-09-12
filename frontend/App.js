import React, { useEffect, useState } from "react";
import { render } from "react-dom";
import { Router } from "@reach/router";
import UserContext from "./UserContext";
import * as Constants from "./constants";
import PrivateRoute from "./PrivateRoute";
import Loading from "./Loading";

const App = () => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Small time out so the loading page doesn't just flash the screen.
    const timer = setTimeout(() => {
      fetch(Constants.API_URL + "/users/me", { credentials: "include" })
        .then(res => res.json())
        .then(result => {
          if (result.status === "success") {
            setUser(true);
          } else {
            setUser(false);
          }
        });
    }, 700);
    return () => clearTimeout(timer);
  }, []);
  return (
    <UserContext.Provider value={user}>
      {user === null ? (
        <Loading />
      ) : (
        <Router>
          <PrivateRoute as={Dashboard} path="/" />
        </Router>
      )}
    </UserContext.Provider>
  );
};
const Dashboard = () => {
  return "Protected Dashboard!";
};

render(<App />, document.getElementById("root"));
