package bootstrap

import (
	"gohub/routers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleware(router)

	// 注册路由
	routers.RegisterAPIRoutes(router)

	// 404
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}

func setup404Handler(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			ctx.String(http.StatusNotFound, "404 页面")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
