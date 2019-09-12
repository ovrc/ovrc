import React from "react";

const Loading = () => {
  return (
    <div id="loading">
      <div id="loading-center">
        <div className="sk-chasing-dots">
          <div className="sk-child sk-dot1"></div>
          <div className="sk-child sk-dot2"></div>
        </div>
      </div>
    </div>
  );
};

export default Loading;
