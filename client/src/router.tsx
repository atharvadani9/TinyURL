import type { RouteObject } from "react-router";
import TinyURL from "./pages";
import Redirect from "./pages/redirect";

const routes: RouteObject[] = [
  {
    path: "/:param",
    element: <Redirect />,
  },
  {
    path: "/404",
    element: <TinyURL />,
  },
  {
    path: "*",
    element: <TinyURL />,
  },
];

export default routes;
