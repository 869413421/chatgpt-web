import serviceAxios from "./request";

export const getUserInfo = () => {
    return serviceAxios({
        url: "auth/info",
        method: "post",
    });
};

export const login = (params: Object) => {
    return serviceAxios({
        url: "user/auth",
        method: "post",
        data: params,
    });
};

export const completion = (chatContext:any) => {
    return serviceAxios({
        url: "chat/completion",
        method: "post",
        data: {
            messages: chatContext,
        },
    });
};

