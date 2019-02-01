package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
// 获取signKey
func GetSignKey() []byte {
	return []byte("doubleduck")
}
func GenerateToken(user *UserName) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),// 可以添加过期时间
	})

	return token.SignedString(GetSignKey())//对应的字符串请自行生成，最后足够使用加密后的字符串
}

func TokenMiddleware()  gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("authorization")
		if tokenStr == "" {
			response(c,http.StatusUnauthorized,"not authorized")
			c.Abort()
			return
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					response(c,http.StatusUnauthorized,"not authorized")
					c.Abort()
					return nil, fmt.Errorf("not authorization")
				}
				return GetSignKey(), nil
			})
			if !token.Valid {
				response(c,http.StatusUnauthorized,"not authorized")
				c.Abort()
				return
			} else {
				c.Next()
				//
				//c.Set("claims", token)
				//return
			}
		}
	}
}

