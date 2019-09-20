import React, { useState, useContext } from "react";
import UserContext from "./UserContext";
import { api_request } from "./Helpers";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loginMessage, setLoginMessage] = useState("");

  const { user, setUser } = useContext(UserContext);

  function login() {
    api_request("/auth/login", "POST", {
      username: username,
      password: password
    }).then(res => {
      if (res.status === "success") {
        setUser(true);
      } else {
        setLoginMessage(res.data.validation);
      }
    });
  }

  return (
    <div id="login-container">
      <div id="login-box" className="box">
        {loginMessage ? (
          <div className="notification is-danger"> {loginMessage} </div>
        ) : null}
        <form
          onSubmit={e => {
            e.preventDefault();
            login();
          }}
        >
          <div className="field">
            <label className="label">Username</label>
            <div className="control">
              <input
                className="input"
                type="text"
                placeholder="Username"
                required
                value={username}
                onChange={e => setUsername(e.target.value)}
              />
            </div>
          </div>
          <div className="field">
            <label className="label">Password</label>
            <div className="control">
              <input
                className="input"
                type="password"
                placeholder="**************"
                required
                value={password}
                onChange={e => setPassword(e.target.value)}
              />
            </div>
          </div>
          <div className="field is-grouped">
            <div className="control">
              <button className="button is-link">Submit</button>
            </div>
            <div className="control">
              <button className="button is-text">Recover Password?</button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
