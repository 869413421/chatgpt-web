import React, { useState } from "react";
import { Flex, FlexItem, Input, Button, toast } from "@chatui/core";
import { useNavigate } from "react-router-dom";
import css from "../../App.module.css";
import "./index.less";
import "@chatui/core/dist/index.css";
import "@chatui/core/es/styles/index.less";
import "md-editor-rt/lib/style.css";
import {login} from '../../services/port'
import { setCookie, getCookie } from "../../utils/cookie";
interface LoginFormState {
  username: string;
  password: string;
}
const Login = () => {
  const navigate = useNavigate();
  const [loginForm, setLoginForm] = useState<LoginFormState>({
    username: "",
    password: "",
  });

  const submitLogin = async () => {
    if (loginForm.username !== "" && loginForm.password !== "") {
      const res = await login(loginForm);
      if (res.data.code === 200) {
        setCookie("mojolicious", res.data.data.token, 3);
        navigate("/");
      } else {
        return toast.show("账号或密码错误", undefined);
      }
    } else {
      return toast.show("请检查账号与密码是否为空", undefined);
    }
  };
  const handleInputChange = (event: any) => {
    const { name, value } = event.target;
    setLoginForm({ ...loginForm, [name]: value });
  };

  return (
    <div className={css.app}>
      <Flex center direction={"column"} style={{ background: "var(--gray-7)" }}>
        <FlexItem
          flex={"1"}
          style={{ marginLeft: "1em" }}
          className="form-Item"
        >
          <div className="login-header">
            <h3>马上开启 AI 之旅</h3>
          </div>
        </FlexItem>
        <FlexItem
          flex={"6"}
          style={{ marginLeft: "1em" }}
          className="form-Item"
        >
          <div className={css.m_top}>
            <input
              className="input-item"
              type="text"
              name="username"
              id="username"
              value={loginForm.username}
              onChange={handleInputChange}
              placeholder="请输入账号"
            />
          </div>
          <div className={css.m_top}>
            <input
              className="input-item"
              type="password"
              name="password"
              id="password"
              value={loginForm.password}
              onChange={handleInputChange}
              placeholder="请输入密码"
            />
          </div>
          <div className={css.m_top}>
            <Button color="primary" block onClick={() => submitLogin()}>
              登陆
            </Button>
          </div>
        </FlexItem>
      </Flex>
    </div>
  );
};

export default Login;
