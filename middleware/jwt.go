package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
	"todo_list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK

		token := c.GetHeader("Authorization")
		
		token = strings.TrimPrefix(token, "Bearer ") // 去掉 "Bearer "，保留真正的 Token
		
		if token == "" {
			code = http.StatusBadRequest
		} else { //解析token
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = http.StatusForbidden //鉴权未通过
				log.Fatal("JWT: ", err)
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = http.StatusUnauthorized //401, Token无效
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    "code解析错误",
			})
			c.Abort()	// 终止请求，不执行后续处理
			return
		}
		c.Next()	//继续执行后续的中间件或 handler。
	}
}
