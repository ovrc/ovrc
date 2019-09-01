import React from "react";
import { render } from "react-dom";
import Login from "./Login";
import AuthContext from "./AuthContext";
import LayoutContext from "./LayoutContext";
import * as Constants from "./constants";
import LoadingScreen from "./LoadingScreen";

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      auth: {
        logged_in: false
      },
      layout: {
        show_loading: false
      }
    };
  }

  componentDidMount() {
    fetch(Constants.API_URL + "/auth/check")
      .then(res => res.json())
      .then(result => {
        if (result.authenticated) {
          this.setState({
            auth: {
              logged_in: true
            }
          });
        }
      });
  }

  render() {
    return (
      <LayoutContext.Provider value={this.state.layout}>
        <LoadingScreen />
        <AuthContext.Provider value={this.state.auth}>
          <Login />
        </AuthContext.Provider>
      </LayoutContext.Provider>
    );
  }
}

render(<App />, document.getElementById("root"));
