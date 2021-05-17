package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	v1 "github.com/xi102/collection/api/v1"
	"github.com/xi102/collection/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//CORS配置
	router.Use(cors.Default())
	router.Use(middleware.LoggerToFile())
	// router.Use(middleware.Cors())

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello,xi102",
	// 	})
	// })

	// router.Static("/index", "./public")
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	router.StaticFS("/files", http.Dir("./upload"))
	router.POST("/upload/", v1.FileUpload)
	return router
}
