package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
)

type receipt struct {
	ID 					int `json:"id,string"`
	Retailer 			string `json:"retailer"`
	PurchaseDate 		string `json:"purchaseDate"`
	PurchaseTime 		string `json:"purchaseTime"`
	Items 				[]item `json:"items,string"`
	Total 				float64 `json:"total,string"`
}

type item struct {
	ShortDescription 	string `json:"shortDescription"`
	Price 				float64 `json:"price,string"`
}

type pointsResponse struct {
	Points int `json:"points"`
}

type processReceiptResponse struct {
	ID int `json:"id"`
}

type allReceipts []receipt

var receipts = allReceipts{
	{
		ID: 1,
		Retailer: "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total: 2.65,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
			{ShortDescription: "Dasani", Price: 1.40},
		},
	},
	{
		ID: 2,
		Retailer: "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total: 1.25,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
		},
	},
	{
		ID: 3,
		Retailer: "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []item{
			{ShortDescription: "Mountain Dew 12PK", Price: 6.49},
			{ShortDescription: "Emils Cheese Pizza", Price: 12.25},
			{ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
			{ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: 12.00},
		},
		Total: 35.35,
	},
	{
		ID: 4,
		Retailer: "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []item{
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
			{ShortDescription: "Gatorade", Price: 2.25},
		},
		Total: 9.00,
	},
}

func main() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")

	fmt.Println("Binding server to port 8080")
	http.ListenAndServe(":8080", r)
}

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

func calculatePoints(receipt receipt) int {
	points := 0

	//One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer{
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	//50 points if the total is a round dollar amount with no cents.
	truncatedTotal := math.Trunc(receipt.Total)
	if truncatedTotal == receipt.Total {
		points += 50
	}

	//25 points if the total is a multiple of 0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	//5 points for every two items on the receipt
	counter := 0.0
	for range receipt.Items {
		counter += 0.5
	}
	counter = math.Trunc(counter)
	points += int(counter) * 5

	//If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. Add the result to points.
	for _, item := range receipt.Items {
		trimmedDesc := strings.Trim(item.ShortDescription, " ")
		if len(trimmedDesc) % 3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd.
	day := strings.Split(receipt.PurchaseDate, "-")[2]
	dayInt, _ := strconv.Atoi(day)
	if dayInt % 2 != 0 {
		points += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm.
	hour := strings.Split(receipt.PurchaseTime, ":")[0]
	minute := strings.Split(receipt.PurchaseTime, ":")[1]
	hourInt, _ := strconv.Atoi(hour)
	minuteInt, _ := strconv.Atoi(minute)
	if (hourInt >= 14 && minuteInt > 00) && hourInt < 17 {
		points += 10
	}

	return points
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

	if(receiptID != -1){
		fmt.Println("Receipt already exists with ID", receiptID, ". Returning that ID.")
	}else{
		fmt.Println("Receipt does not exist. Adding receipt to receipts array. Returning with ID", len(receipts) + 1)
		receipt.ID = len(receipts) + 1
		receipts = append(receipts, receipt)
		receiptID = receipt.ID
	}
	
	response := processReceiptResponse{ID: receiptID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println("Successfully processed receipt ID", receiptID)
}


func findReceiptID(receipt receipt) int {
	for _, r := range receipts {
		if r.Retailer == receipt.Retailer && r.PurchaseDate == receipt.PurchaseDate && r.PurchaseTime == receipt.PurchaseTime && r.Total == receipt.Total && len(r.Items) == len(receipt.Items) {
			return r.ID
		}
	}

	return -1
}