import React, { useState } from "react";
import { api_request } from "./Helpers";

const Users = () => {
  const [users, setUsers] = useState(null);

  api_request("/users/list", "GET").then(res => {
    if (res.status === "success") {
      setUsers(true);
    } else {
      setUsers(false);
    }
  });

  if (users === null) {
    return "Loading...";
  }

  if (users === true) {
    return "User list!";
  }

  return "Users!";
};

export default Users;
