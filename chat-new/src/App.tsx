import { RouterProvider } from 'react-router-dom';
import { routes } from './routers';

const App = () => {
  return <RouterProvider router={routes} />
}

export default App
