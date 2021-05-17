package v1

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	// name := c.PostForm("name")
	// email := c.PostForm("email")

	// Multipart form
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
	// 	return
	// }
	// files := form.File["files"]

	// for _, file := range files {
	// 	filename := filepath.Base(file.Filename)
	// 	if err := c.SaveUploadedFile(file, filepath.Join("upload", filename)); err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
	// 		return
	// 	}
	// }

	// c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email))
	file, err := c.FormFile("file")
	log.Println(file.Filename)

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filepath.Join("./upload", filename)); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	log.Println("file\t" + filename + "saved sucessfully")
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))

}
