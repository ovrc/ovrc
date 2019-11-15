import React, { useState, useEffect } from "react";
import { api_request } from "./Helpers";

const Users = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    api_request("/users", "GET").then(res => {
      if (res.status === "success") {
        setUsers(true);
      } else {
        setUsers(false);
      }
    });
  }, []);

  if (users === true) {
    return "User list!";
  }

  return "Loading...";
};

export default Users;
