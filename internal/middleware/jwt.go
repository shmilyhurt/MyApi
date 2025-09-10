package middleware

import (
	"MyApi/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// JWTAuthMiddleware 验证 JWT 的中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 中取出 token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少 token"})
			c.Abort()
			return
		}

		// Bearer token 格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token 格式错误"})
			c.Abort()
			return
		}

		// 解析 token
		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效或过期的 token"})
			c.Abort()
			return
		}

		// 把用户信息存到 context
		c.Set("userID", claims.UserID)

		// 继续执行
		c.Next()
	}
}
