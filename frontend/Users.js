import React, { useState, useEffect } from "react";
import { api_request } from "./Helpers";

const Users = () => {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    api_request("/users", "GET").then(res => {
      if (res.status === "success") {
        setUsers(res.data.users);
      } else {
        setUsers(false);
      }
    });
  }, []);

  if (users.length > 0) {
    return (
      <table className="table">
        <thead>
          <tr>
            <th>Username</th>
            <th>Dt. Created</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map(function(user) {
            return (
              <tr>
                <td>{user.username}</td>
                <td>{user.dt_created}</td>
                <td>
                  <span className="icon has-text-info">
                    <i className="fas fa-edit"></i>
                  </span>
                  <span className="icon has-text-danger">
                    <i className="fas fa-ban"></i>
                  </span>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    );
  }

  return "Loading...";
};

export default Users;
