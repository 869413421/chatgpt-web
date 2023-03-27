package controllers

import (
	"github.com/869413421/chatgpt-web/pkg/auth"
	"github.com/869413421/chatgpt-web/pkg/model/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// AuthController 认证控制器
type AuthController struct {
	BaseController
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// authRequest 认证请求
type authRequest struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

// Auth 认证
func (c *AuthController) Auth(ctx *gin.Context) {
	var req authRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if req.Name == "" || req.Password == "" {
		c.ResponseJson(ctx, http.StatusUnauthorized, "请输入用户名密码", nil)
		return
	}

	authUser, err := user.GetByName(req.Name)
	if err != nil && err == gorm.ErrRecordNotFound {
		c.ResponseJson(ctx, http.StatusUnauthorized, "请求认证的用户不存在", nil)
		return
	}
	if !authUser.ComparePassword(req.Password) {
		c.ResponseJson(ctx, http.StatusUnauthorized, "密码错误", nil)
		return
	}
	token, err := auth.Encode(authUser)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	c.ResponseJson(ctx, http.StatusOK, "", gin.H{
		"token": token,
	})
}

// Info 登录用户信息
func (c *AuthController) Info(ctx *gin.Context) {
	authUser, ok := ctx.Get("authUser")
	if !ok {
		c.ResponseJson(ctx, http.StatusInternalServerError, "获取登录用户信息失败", nil)
		return
	}

	userInfo, ok := authUser.(*user.User)
	if !ok {
		c.ResponseJson(ctx, http.StatusInternalServerError, "断言登录用户信息失败", nil)
		return
	}
	// 未实现权限系统，写死
	c.ResponseJson(ctx, http.StatusOK, "", gin.H{
		"info":             userInfo,
		"permissionRoutes": []string{"chat", "chat/completion", "user/auth/info", "user/auth"},
	})
}
