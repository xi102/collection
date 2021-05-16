package routers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello,xi102",
		})
	})

	// router.Static("/index", "./public")
	router.Use(static.Serve("/index", static.LocalFile("./frontend/dist", false)))
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	router.StaticFS("/files", http.Dir("./upload"))
	router.POST("/upload/", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filepath.Join("upload", filename)); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}

		c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	})
	return router
}
