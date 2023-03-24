package auth

import (
	"errors"
	"github.com/869413421/chatgpt-web/pkg/model/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var (
	key = []byte("pgServiceUserTokenKeySecret")
)

type CustomClaims struct {
	User *user.User
	jwt.StandardClaims
}

// Decode a token string into a token object
func Decode(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func Encode(user *user.User) (string, error) {

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "chatgpt-web",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// EncodeByCtx 从ctx中的token获取登录用户信息
func EncodeByCtx(c *gin.Context) (*CustomClaims, error) {
	//1.获取token
	token := c.GetHeader("Authorization")
	if token != "" {
		tokenS := strings.Split(token, " ")
		token = tokenS[1]
	} else {
		token = c.Request.FormValue("token")
	}
	if token == "" {
		return nil, errors.New("not found token")
	}

	return Decode(token)
}
