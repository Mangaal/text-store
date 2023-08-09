package testapis_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"textstore/pkg/apis"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFileOption(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/apis/file/option/:sort/:limit", apis.DeleteFile)

	req := httptest.NewRequest("GET", "/apis/file/option/a/10", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestGetFiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/apis/files", apis.GetFiles)

	req := httptest.NewRequest("GET", "/apis/files", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
