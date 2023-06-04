package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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