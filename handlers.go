package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	fmt.Printf("Received request to get points for receipt ID %v \n", id)

	points := 0

	if err != nil {
		fmt.Println("Error converting ID in reqeust to type int")
	}

	//loop through all items in the allReceipts array and find the one with the given id
	for _, item := range receipts {
		if item.ID == id {
			points = calculatePoints(item)
		}
	}

	response := pointsResponse{Points: points}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Printf("Returning %d points for receipt ID %v \n", points, id)
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request to process receipt.")

	// Convert JSON request body to receipt struct
	var receipt receipt

	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if receipts array has an identical receipt to the one in the request body
	// If it does, return the ID of the receipt in the receipts array
	// If it doesn't, add an ID of len(receipts) + 1 to the receipt and append it to the receipts array, then return the new ID
	receiptID := findReceiptID(receipt)

	if receiptID != -1 {
		fmt.Println("Receipt already exists with ID", receiptID, ". Returning that ID.")
	} else {
		fmt.Println("Receipt does not exist. Adding receipt to receipts array. Returning with ID", len(receipts)+1)
		receipt.ID = len(receipts) + 1
		receipts = append(receipts, receipt)
		receiptID = receipt.ID
	}

	response := processReceiptResponse{ID: receiptID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println("Successfully processed receipt ID", receiptID)
}