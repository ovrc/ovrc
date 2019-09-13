import * as Constants from "./constants";

export const api_request = endpoint => {
  // Function to make requests to the backend API.
  return fetch(Constants.API_URL + endpoint, { credentials: "include" }).then(
    res => {
      if (res.status < 500) {
        return res.json();
      }
    }
  );
};
