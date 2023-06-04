package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting server")
	r := mux.NewRouter()

	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")

	fmt.Println("Binding server to port 8080")
	http.ListenAndServe(":8080", r)
}