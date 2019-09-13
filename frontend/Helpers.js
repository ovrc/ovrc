import * as Constants from "./constants";
import axios from "axios";
import qs from "qs";

export const api_request = (endpoint, method, params) => {
  if (!method) {
    method = "GET";
  }

  if (!params) {
    params = {};
  }

  const url = Constants.API_URL + endpoint;

  return axios({
    method: method,
    url: url,
    withCredentials: true,
    headers:
      method === "POST"
        ? { "content-type": "application/x-www-form-urlencoded" }
        : null,
    params: method === "GET" && params ? params : null,
    data: method === "POST" && params ? qs.stringify(params) : null,
    validateStatus: function(status) {
      return status >= 200 && status < 500;
    }
  })
    .then(res => {
      return res.data;
    })
    .catch(function(error) {
      console.log(error);
    });
};
