import React, { useContext } from "react";
import UserContext from "./UserContext";
import Login from "./Login";

const PrivateRoute = props => {
  const { user, _ } = useContext(UserContext);
  let { as: Comp } = props;

  return user ? <Comp {...props} /> : <Login />;
};

export default PrivateRoute;
