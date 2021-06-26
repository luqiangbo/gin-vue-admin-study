package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goClass/model/request"
)

func jWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt健全取头部信息 x-token 登录时返回token信息 这里前端需要吧token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录

	}
}

type JWT struct {
	SigningKey []byte
}

// 创建一个token

func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
