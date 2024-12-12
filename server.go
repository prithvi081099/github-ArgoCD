package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Response struct for sending JSON responses
type Response struct {
	Sum int `json:"sum"`
}

// Handler function for the API
func GetExampleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input1Str := vars["input1"]
	input2Str := vars["input2"]

	// Convert inputs to integers
	input1, err1 := strconv.Atoi(input1Str)
	input2, err2 := strconv.Atoi(input2Str)

	// Handle conversion errors
	if err1 != nil || err2 != nil {
		http.Error(w, "Inputs must be valid integers", http.StatusBadRequest)
		return
	}

	// Perform some logic (e.g., calculate sum)
	sum := input1 + input2

	// Create response
	response := Response{
		Sum: sum,
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func main() {
	// Create a new router
	router := mux.NewRouter()

	// Define the API route
	router.HandleFunc("/add/{input1}/{input2}", GetExampleHandler).Methods("GET")

	// Start the server
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
