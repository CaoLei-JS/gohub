package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Hello":"World!",
			})
		})
	}
}
