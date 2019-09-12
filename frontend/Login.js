import React from "react";
import UserContext from "./UserContext";

const Login = () => {
  return (
    <UserContext.Consumer>
      {(user, setUser) => <div></div>}
    </UserContext.Consumer>
  );
};

export default Login;
