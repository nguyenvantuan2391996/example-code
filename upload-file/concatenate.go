package main

import (
	"fmt"
	"io"
	"os"
)

func concatenateVideo(length int) bool {
	// Open the output file for writing
	outputFile, err := os.Create("./upload-file/output.mp4")
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer func(outputFile *os.File) {
		errClose := outputFile.Close()
		if errClose != nil {
			return
		}
	}(outputFile)

	// List the video chunks in the desired order
	chunkPaths := make([]string, 0)
	for i := 0; i < length; i++ {
		chunkPaths = append(chunkPaths, fmt.Sprintf("upload-file/video%v", i))
	}

	// Concatenate the video chunks
	for _, chunkPath := range chunkPaths {
		// Open the current chunk file
		chunkFile, errChunk := os.Open(chunkPath)
		if errChunk != nil {
			fmt.Println("Error:", errChunk)
			return false
		}
		defer func(chunkFile *os.File) {
			errClose := chunkFile.Close()
			if errClose != nil {
				return
			}
		}(chunkFile)

		// Copy the chunk content to the output file
		_, errCopy := io.Copy(outputFile, chunkFile)
		if errCopy != nil {
			fmt.Println("Error:", errCopy)
			return false
		}
	}

	return true
}

func main() {
	isConcatenate := concatenateVideo(8)
	if !isConcatenate {
		fmt.Println("concatenate is failed")
		return
	}

	fmt.Println("concatenate is successfully")
}
