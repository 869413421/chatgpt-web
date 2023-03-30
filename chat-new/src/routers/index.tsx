import { lazy, Suspense } from "react";
import {createBrowserRouter, createMemoryRouter, Navigate, redirect} from "react-router-dom";
import type { RouteObject } from "react-router-dom";
import ErrorBoundary from "../components/ErrorBoundary";
import { getUserInfo } from "../services/port";
// 不需要懒加载的页面组件
import Permission from "../components/Permission";
import NoFind from "../pages/noFind";
// 需要懒加载的页面组件
const ChatContainer = lazy(() => import("../pages/chat"));
const Login = lazy(() => import("../pages/login"));
/**
 * @param Component 懒加载的组件
 * @param code 用于判断权限的字段(你可以自己定)
 * @returns
 */
const LazyLoad = (
  Component: React.LazyExoticComponent<() => JSX.Element>,
  code?: string
) => {
  return (
    <Permission code={code}>
      <Suspense fallback={<div style={{ margin: "auto" }}>loading...</div>}>
        <Component />
      </Suspense>
    </Permission>
  );
};

export interface UserInfo {
  name: string;
  permissionRoutes: string[];
  code: number;
}

/**
 * @description 这个loader函数会在路由渲染前触发,所以可以用来做路由权限控制和登陆重定向
 * @description (取代请求拦截器中的登陆重定向)
 * @description 这个loader函数返回值可以在页面中通过 useRouteLoaderData(id)或者useLoaderData获取
 */
const rootLoader = async () => {
  // console.log('页面加载前请求用户信息')
  // 这里用假的接口模拟下
  const res = await getUserInfo();
  if (res.status == 401){
    return redirect("/login");
  }
  const { info, permissionRoutes } = res.data.data;
  return {
    info,
    permissionRoutes,
  };
};

const routerConfig: RouteObject[] = [
  {
    path: "/",
    element: <Navigate to="/chat" />,
  },
  {
    path: "/chat",
    id: "root",
    errorElement: <ErrorBoundary />,
    loader: rootLoader,
    element: LazyLoad(ChatContainer, "chat"),
  },
  {
    path: "/login",
    element: LazyLoad(Login),
  },
  {
    path: "*",
    element: <NoFind />,
  },
];

export const routes = createMemoryRouter(routerConfig);
