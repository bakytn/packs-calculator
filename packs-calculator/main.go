package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"repartners/packs-calculator/packs"
	"sort"
)

// Request represents the JSON structure for incoming requests
type Request struct {
	ItemsOrdered int    `json:"items_ordered"`
	PackSizes    string `json:"pack_sizes,omitempty"` // Optional: comma-separated pack sizes
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

	// Parse the JSON request
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.ItemsOrdered <= 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Default pack sizes
	packSizesStr := req.PackSizes
	if packSizesStr == "" {
		packSizesStr = "250,500,1000,2000,5000"
	}

	// Parse pack sizes
	packSizesList, err := packs.ParsePackSizes(packSizesStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing pack sizes: %v", err), http.StatusBadRequest)
		return
	}

	// Calculate the required packs
	result := packs.CalculatePacks(req.ItemsOrdered, packSizesList)

	// Prepare the response
	var packsToSend []packs.PackCount
	for packSize, count := range result {
		packsToSend = append(packsToSend, packs.PackCount{PackSize: packSize, Count: count})
	}

	// Sort the response for consistency
	sort.Slice(packsToSend, func(i, j int) bool {
		return packsToSend[i].PackSize < packsToSend[j].PackSize
	})

	response := Response{
		ItemsOrdered: req.ItemsOrdered,
		PacksToSend:  packsToSend,
	}

	// Send the JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
