import { useRouteError } from "react-router-dom";

const ErrorBoundary = () => {
   const err = useRouteError() as any;
   return <div>
      <p>
         出错啦～
      </p>
      <p>
         错误信息: {err.message}
      </p>
   </div>
}

export default ErrorBoundary;