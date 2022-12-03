package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
	r.NoRoute(func(context *gin.Context) {
		acceptString := context.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html"){
			context.String(http.StatusNotFound,"页面返回404")
		}else {
			context.JSON(http.StatusNotFound,gin.H{
				"error_code":404,
				"error_message":"路由未定义",
			})
		}
	})
	r.Run(":8000")
}
