package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DeleteFile(c *gin.Context) {

	filebody := FileBody{}

	body, err := c.GetRawData()

	if err != nil {

		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	json.Unmarshal(body, &filebody)

	for _, fileName := range filebody.Files {

		filePath := uploadDirectory + "/" + fileName

		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error deleting file:", err)
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return

		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Files Deleted successfully"})

}

func UpdateFileContent(c *gin.Context) {

	newname := c.Param("newname")

	oldname := c.Param("oldname")

	fmt.Println("new Name", newname)

	err := os.Rename(uploadDirectory+"/"+oldname, uploadDirectory+"/"+newname)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Files Updated successfully"})

}
