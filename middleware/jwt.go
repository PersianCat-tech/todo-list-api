package middleware

import (
	"net/http"
	"time"
	"todo_list/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusBadRequest
		} else { //解析token
			claim, err := utils.ParseToken(token)
			if err != nil {
				code = http.StatusForbidden //鉴权未通过
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
