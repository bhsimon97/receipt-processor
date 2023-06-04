package main

import (
	"fmt"
	"testing"
)

func TestCalculatePoints(t *testing.T){
	fmt.Println("Testing calculatePoints()")
	receipt := receipt{
		ID: 2,
		Retailer: "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total: 1.25,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
		},
	}

	expectedPoints := 31
	calculatedPoints := calculatePoints(receipt)

	if calculatedPoints != expectedPoints {
		t.Errorf("Expected points: %d, but got: %d", expectedPoints, calculatedPoints)
	}
}

func TestFindReceiptID_MatchFound(t *testing.T){
	fmt.Println("Testing findReceiptID() - Match Found Case")
	receipt := receipt{
		ID: 1,
		Retailer: "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total: 2.65,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
			{ShortDescription: "Dasani", Price: 1.40},
		},
	}

	expectedID := 1
	foundID := findReceiptID(receipt)
	
	if foundID != expectedID {
		t.Errorf("Expected receipt ID: %d, but got: %d", expectedID, foundID)
	}
}

func TestFindReceiptID_MatchNotFound(t *testing.T){
	fmt.Println("Testing findReceiptID() - Match Not Found Case")
	receipt := receipt{
		Retailer: "Costco",
		PurchaseDate: "2022-06-02",
		PurchaseTime: "12:13",
		Total: 2.65,
		Items: []item{
			{ShortDescription: "Pepsi - 12-oz", Price: 1.25},
			{ShortDescription: "Dasani", Price: 1.40},
		},
	}

	expectedID := -1
	foundID := findReceiptID(receipt)

	if foundID != expectedID {
		t.Errorf("Expected receipt ID: %d, but got: %d", expectedID, foundID)
	}
}