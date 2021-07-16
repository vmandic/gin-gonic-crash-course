package main

import (
	"io"
	"net/http"
	"os"

	"github.com/vmandic/gin-gonic-crash-course/controller"
	"github.com/vmandic/gin-gonic-crash-course/middlewares"
	"github.com/vmandic/gin-gonic-crash-course/service"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setLogOutput()

	server := gin.New()
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("./templates/*.html")

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	server.Use(middlewares.BasicAuth())
	server.Use(gindump.Dump())

	apiGroup := server.Group("/api")
	{
		apiGroup.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})

		apiGroup.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiGroup.POST("/posts", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(200, gin.H{"message": "Input valid."})
			}
		})

		viewRoutes := server.Group("/view")
		{
			viewRoutes.GET("/videos", videoController.ShowAll)
		}
	}

	server.Run(":8080")
}
