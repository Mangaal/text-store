package main

import (
	"net/http"
	"textstore/pkg/apis"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi You have successfully Install STORE APi",
		})
	})

	r.GET("apis/files", apis.GetFiles)

	r.POST("apis/file", apis.StoreFile)

	r.GET("apis/file/option/:sort/:limit", apis.FileOption)

	r.POST("apis/file/:newname/:oldname", apis.UpdateFileContent)

	r.DELETE("apis/file", apis.DeleteFile)

	r.Run(":80")

}
