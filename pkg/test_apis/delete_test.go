package testapis_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"textstore/pkg/apis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type FileBody struct {
	Files []string `json:"files"`
}

func TestDeleteFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/apis/file", apis.DeleteFile)
	file := FileBody{}

	file.Files = append(file.Files, "newfile.txt")

	body, _ := json.Marshal(file)
	req := httptest.NewRequest("POST", "/apis/file", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}
