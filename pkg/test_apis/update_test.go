package testapis_test

import (
	"fmt"
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

func TestUpdateFileContent(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/apis/file/:newname/:oldname", apis.UpdateFileContent)

	oldName := "oldfile.txt"
	newName := "newfile.txt"

	req := httptest.NewRequest("POST", fmt.Sprintf("/apis/file/%s/%s", newName, oldName), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}
