import { useRoutes, BrowserRouter } from "react-router-dom";

import routes from "routes.js";

function App() {
  const routing = useRoutes(routes);
  return (
    <>
      <div className="App">{routing}</div>
    </>
  );
}

export default App;
