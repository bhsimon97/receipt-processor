package handlers

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