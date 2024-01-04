package main

import (
	"log"
	"net/http"
)

func generateShortUrlHandler(w http.ResponseWriter, r *http.Request) {

}

func getLongUrlHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/generate-short-url", generateShortUrlHandler)
	http.HandleFunc("/get-long-url", getLongUrlHandler)

	// Start the server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
