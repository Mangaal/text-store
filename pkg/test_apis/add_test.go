package testapis_test

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
	"textstore/pkg/apis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStoreFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/apis/file", apis.StoreFile)

	files := []struct {
		FileName    string
		FileContent string
	}{
		{FileName: "oldfile.txt", FileContent: "Today I am learning new things"},

		// Add more files as needed
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, file := range files {
		part, _ := writer.CreateFormFile("files", file.FileName)
		part.Write([]byte(file.FileContent))
	}
	writer.Close()

	req := httptest.NewRequest("POST", "/apis/file", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestStoreFile_NoFilesAttached(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/apis/file", apis.StoreFile)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	req := httptest.NewRequest("POST", "/apis/file", nil)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
