package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductName string `json:"product_name"`
	ID          int    `json:"id"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode([]*Product{
		{
			ID:          1,
			ProductName: "iPhone 16 Pro Max",
		},
		{
			ID:          2,
			ProductName: "iPhone 8 Plus",
		},
	})
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/products", getProducts)

	// Start the server
	log.Fatal(http.ListenAndServe(":3001", nil))
}
