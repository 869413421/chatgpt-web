import serviceAxios from "./request";

export const getUserInfo = () => {
  return serviceAxios({
    url: "/api/user/info",
    method: "get",
  });
};

export const login = (params: Object) => {
  return serviceAxios({
    url: "/api/user/login",
    method: "post",
    params,
  });
};

