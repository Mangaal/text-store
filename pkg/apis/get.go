package apis

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetFiles(c *gin.Context) {

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

func FileOption(c *gin.Context) {

	Count := 0

	limit := c.Param("limit")
	order := c.Param("sort")

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
