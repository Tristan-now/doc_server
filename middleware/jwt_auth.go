package middleware

import (
	"github.com/gin-gonic/gin"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
)

// JwtAuth 验证用户是否登录
func JwtAuth() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		ok := redis_service.CheckLogout(token)
		if ok {
			res.FailWithMsg("token已注销", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
