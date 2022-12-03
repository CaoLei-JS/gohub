package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/routes"
	"net/http"
	"strings"
)

func SetRoute(router *gin.Engine) {
	registerGlobalMiddleware(router)
	routes.RegisterApiRoutes(router)
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(context *gin.Context) {
		acceptString := context.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			context.String(http.StatusNotFound, "页面返回404")
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "未定义",
			})
		}

	})
}
