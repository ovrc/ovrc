import React from "react";
import { render } from "react-dom";
import Login from "./Login";
import { Provider } from "./AuthContext";
import * as Constants from "./constants";

var loading_container = document.getElementById("loading");

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      logged_in: false
    };
  }

  componentDidMount() {
    fetch(Constants.API_URL + "/auth/check")
      .then(res => res.json())
      .then(
        result => {
          if (result.authenticated) {
            this.setState({
              logged_in: true
            });
          }

          loading_container.classList.toggle("fade");
        },
        error => {}
      );
  }

  render() {
    return (
      <Provider value={this.state}>
        {this.state.logged_in ? console.log("logged in") : <Login />}
      </Provider>
    );
  }
}

render(<App />, document.getElementById("root"));
