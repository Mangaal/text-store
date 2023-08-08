package apis

import (
	"os"
)

type FileBody struct {
	Files []string `json:"files"`
}

var uploadDirectory = "uploads"

func uploadDirectoryExists(uploadDirectory string) error {
	err := os.MkdirAll(uploadDirectory, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
