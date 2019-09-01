import React from "react";
import LayoutContext from "./LayoutContext";

class LoadingScreen extends React.Component {
  componentDidMount() {
    console.log("loading screen mounted");
  }

  render() {
    return (
      <LayoutContext.Consumer>
        {layout =>
          layout.show_loading ? (
            <div id="loading">
              <div id="loading-center">
                <div className="sk-chasing-dots">
                  <div className="sk-child sk-dot1"></div>
                  <div className="sk-child sk-dot2"></div>
                </div>
              </div>
            </div>
          ) : null
        }
      </LayoutContext.Consumer>
    );
  }
}

export default LoadingScreen;
