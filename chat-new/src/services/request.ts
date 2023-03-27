import axios from "axios";
import {getCookie} from "../utils/cookie";

const serviceAxios = axios.create({
    withCredentials: false, // 跨域请求是否需要携带 cookie
    baseURL: "http://localhost:8080" // 测试用
});
serviceAxios.interceptors.request.use(
    (config) => {
        if (getCookie("mojolicious")) {
            config.headers["Authorization"] = "Bearer "+getCookie("mojolicious"); // 请求头携带 token
        }
        return config;
    },
    (error) => {
        Promise.reject(error);
    }
);

serviceAxios.interceptors.response.use(
    (res) => {
        return res;
    },
    (error) => {
        let message = "";
        if (error && error.response) {
            switch (error.response.status) {
                case 302:
                    message = "接口重定向了！";
                    break;
                case 400:
                    message = "参数不正确！";
                    break;
                case 401:
                    return error.response
                    // message = "您未登录，或者登录已经超时，请先登录！";
                    break;
                case 403:
                    message = "您没有权限操作！";
                    break;
                case 404:
                    message = `请求地址出错: ${error.response.config.url}`;
                    break;
                case 500:
                    message = "服务器内部错误！";
                    break;
                case 501:
                    message = "服务未实现！";
                    break;
                case 502:
                    message = "网关错误！";
                    break;
                case 503:
                    message = "服务不可用！";
                    break;
                case 504:
                    message = "服务暂时无法访问，请稍后再试！";
                    break;
                case 505:
                    message = "HTTP 版本不受支持！";
                    break;
                default:
                    message = "异常问题，请联系管理员！";
                    break;
            }
        }
        return Promise.reject(message);
    }
);

export default serviceAxios;
