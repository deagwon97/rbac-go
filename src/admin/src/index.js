import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter } from "react-router-dom";
// import { CookiesProvider } from 'react-cookie';
import App from "App";
import reportWebVitals from "reportWebVitals";

ReactDOM.render(
  <BrowserRouter basename="/admin">
    <App />
  </BrowserRouter>,
  document.getElementById("root"),
);

reportWebVitals();
