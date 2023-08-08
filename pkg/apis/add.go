package apis

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StoreFile(c *gin.Context) {

	fmt.Println("receive request for file add")

	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {

		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	files := c.Request.MultipartForm.File["files"]

	if len(files) == 0 {
		fmt.Println("No files attached")

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": errors.New("no file attached")})
		return
	}

	uploadDirectoryExists(uploadDirectory)

	for _, file := range files {
		uploadedFile, err := file.Open()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer uploadedFile.Close()

		// Create a new file on the server to save the uploaded content
		path := filepath.Join(uploadDirectory, file.Filename)

		fmt.Println("creating   ", path)
		serverFile, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer serverFile.Close()

		// Copy the content from the uploaded file to the server file
		_, err = io.Copy(serverFile, uploadedFile)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Files uploaded successfully"})

}
