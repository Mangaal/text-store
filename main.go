package main

import (
	"textstore/pkg/apis"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("apis/files", apis.GetFiles)

	r.POST("apis/file", apis.StoreFile)

	r.GET("apis/file/option/:sort/:limit", apis.FileOption)

	r.POST("apis/file/:newname/:oldname", apis.UpdateFileContent)

	r.DELETE("apis/file", apis.DeleteFile)

	r.Run(":80")

}
