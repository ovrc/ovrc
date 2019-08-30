import React from "react";

const AuthContext = React.createContext({
  logged_in: false,
  handleLoginChange() {}
});

export const Provider = AuthContext.Provider;
export const Consumer = AuthContext.Consumer;
