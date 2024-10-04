package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"repartners/packs-calculator/packs"
)

// Request represents the JSON structure for incoming requests
type Request struct {
	ItemsOrdered int `json:"items_ordered"`
}

// Response represents the JSON structure for responses
type Response struct {
	ItemsOrdered int               `json:"items_ordered"`
	PacksToSend  []packs.PackCount `json:"packs_to_send"`
}

func main() {
	http.HandleFunc("/calculate-packs", calculatePacksHandler)
	fmt.Println("Server is running on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

// HTTP handler for calculating packs
func calculatePacksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.ItemsOrdered <= 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := Response{
		ItemsOrdered: req.ItemsOrdered,
		PacksToSend:  packsToSend,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
