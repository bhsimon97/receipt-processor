package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetPoints(t *testing.T) {
	fmt.Println("Testing getPoints()")
	request := httptest.NewRequest(http.MethodGet, "/receipts/3/points", nil)

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", getPoints)

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}

	var pointsRes pointsResponse
	err := json.NewDecoder(response.Body).Decode(&pointsRes)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Check the points value in the response
	expectedPoints := 28
	if pointsRes.Points != expectedPoints {
		t.Errorf("Expected points %d, but got %d", expectedPoints, pointsRes.Points)
	}
}

func TestProcessReceipt(t *testing.T) {
	fmt.Println("Testing processReceipt()")
	expectedID := len(receipts) + 1 

	receipt := receipt{
		Retailer:      "Walgreens",
		PurchaseDate:  "2022-01-17",
		PurchaseTime:  "12:25",
		Total:         2.65,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
			{ShortDescription: "Dasani", Price: 1.40},
		},
	}

	body, _ := json.Marshal(receipt)

	// Create a request with the JSON body
	request := httptest.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(body))

	// Create a ResponseRecorder to capture the response
	response := httptest.NewRecorder()

	// Create a new router and register the handler function
	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", processReceipt)

	// Serve the request
	router.ServeHTTP(response, request)

	// Check the response status code
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.Code)
	}

	// Decode the JSON response
	var processRes processReceiptResponse
	err := json.NewDecoder(response.Body).Decode(&processRes)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	// Check the ID value in the response
	
	if processRes.ID != expectedID {
		t.Errorf("Expected ID %d, but got %d", expectedID, processRes.ID)
	}
}