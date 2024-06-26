package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-Jarvis/admin/rbac"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/global"
)

func CasbinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := global.User
		obj := c.Request.URL.Path
		act := c.Request.Method
		e := rbac.NewCasbin(global.DB) // 判断策略中是否存在
		ok, err := e.Enforce(sub, obj, act)
		if ok {
			c.Next()
		} else {
			global.Logger.Errorf(err.Error())
			response.AuthForbidden(c)
			c.Abort()
		}
	}
}
