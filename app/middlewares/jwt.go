package middlewares

import (
	"github.com/869413421/chatgpt-web/app/http/controllers"
	"github.com/869413421/chatgpt-web/pkg/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var base = controllers.BaseController{}

// Jwt jwt认证
func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := auth.EncodeByCtx(c)
		if err != nil {
			base.ResponseJson(c, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		if claims.User.ID == 0 {
			base.ResponseJson(c, http.StatusUnauthorized, "用户信息错误，未知的token", nil)
			return
		}
		c.Set("authUser", claims.User)
		c.Next()
	}
}
