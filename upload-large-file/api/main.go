package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	index := r.FormValue("chunkIndex")
	file, _, err := r.FormFile("chunk")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		errClose := file.Close()
		if errClose != nil {
			return
		}
	}(file)

	// Process and save the chunk to a file
	outputFile, err := os.Create(fmt.Sprintf("./upload-large-file/video%v", index))
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer func(outputFile *os.File) {
		errClose := outputFile.Close()
		if errClose != nil {
			return
		}
	}(outputFile)

	_, err = io.Copy(outputFile, file)
	if err != nil {
		http.Error(w, "Failed to write file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Chunk uploaded successfully"))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
