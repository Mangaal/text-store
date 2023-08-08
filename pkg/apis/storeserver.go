package apis

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileBody struct {
	Files []string `json:"files"`
}

func FileOption(c *gin.Context) {

	Count := 0

	limit := c.Param("limit")
	order := c.Param("sort")

	uploadDirectory := os.Getenv("API_DIR")
	wordFrequency := make(map[string]int)

	// Open the directory
	dir, err := os.Open(uploadDirectory)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(-1) // -1 means read all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() { // Check if it's a regular file

			file, err := os.Open(uploadDirectory + "/" + fileInfo.Name())
			if err != nil {
				fmt.Println("Error opening file:", err)
				continue
			}
			defer file.Close()

			// Read file content and update word frequency
			reader := bufio.NewReader(file)
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Println("Error reading file:", err)
					break
				}
				words := strings.Fields(line)
				for _, word := range words {
					Count++
					wordFrequency[word]++
				}
			}

		}
	}

	// Create a slice of WordFrequency objects
	var WordFrequencyList []struct {
		Word      string
		Frequency int
	}
	for word, frequency := range wordFrequency {
		WordFrequencyList = append(WordFrequencyList, struct {
			Word      string
			Frequency int
		}{word, frequency})
	}

	if order == "a" {
		// Sort the slice in descending order based on frequency
		sort.SliceStable(WordFrequencyList, func(i, j int) bool {
			return WordFrequencyList[i].Frequency < WordFrequencyList[j].Frequency
		})

	} else {
		// Sort the slice in descending order based on frequency
		sort.SliceStable(WordFrequencyList, func(i, j int) bool {
			return WordFrequencyList[i].Frequency > WordFrequencyList[j].Frequency
		})

	}

	nolimit, _ := strconv.Atoi(limit)

	if len(WordFrequencyList) < nolimit {

		nolimit = len(WordFrequencyList)

	}

	for i, wf := range WordFrequencyList[:nolimit] {
		fmt.Printf("%d. %s (%d occurrences)\n", i+1, wf.Word, wf.Frequency)
	}

	c.JSON(http.StatusOK, gin.H{
		"items":          WordFrequencyList[:nolimit],
		"totalWordCount": Count,
	})

}

func DeleteFile(c *gin.Context) {

	uploadDirectory := os.Getenv("API_DIR")

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

func GetFiles(c *gin.Context) {

	uploadDirectory := os.Getenv("API_DIR")

	Files := []string{}

	// Open the directory
	dir, err := os.Open(uploadDirectory)
	if err != nil {
		fmt.Println("Error opening directory:", err)

		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}
	defer dir.Close()

	// Read the directory entries
	fileInfos, err := dir.Readdir(-1) // -1 means read all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Print file names
	for _, fileInfo := range fileInfos {
		if fileInfo.Mode().IsRegular() { // Check if it's a regular file
			Files = append(Files, fileInfo.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{"files": Files})

}

func UpdateFileContent(c *gin.Context) {

	uploadDirectory := os.Getenv("API_DIR")

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

}

func StoreFile(c *gin.Context) {

	uploadDirectory := os.Getenv("API_DIR")

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

func uploadDirectoryExists(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
