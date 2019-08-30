import React from "react";
import { render } from "react-dom";
import Login from "./Login";
import { Provider } from "./AuthContext";

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      logged_in: false
    };

    fetch("https://127.0.0.1:8002/auth/check")
      .then(res => res.json())
      .then(
        result => {
          if (result.authenticated) {
            this.setState({
              logged_in: true
            });
          }
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
