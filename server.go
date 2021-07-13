package main

import (
	"github.com/vmandic/gin-gonic-crash-course/controller"
	"github.com/vmandic/gin-gonic-crash-course/middlewares"
	"github.com/vmandic/gin-gonic-crash-course/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
