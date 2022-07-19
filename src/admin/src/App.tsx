import { useRoutes } from 'react-router-dom';
import routes from 'routes';

export const API_URL = process.env.REACT_APP_API_URL;

function App() {
  const routing = useRoutes(routes);
  return <div className="App">{routing}</div>;
}

export default App;
