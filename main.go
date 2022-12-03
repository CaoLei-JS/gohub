package main

import (
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetRoute(r)
	r.Run(":8000")
}
