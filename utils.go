package main

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func calculatePoints(receipt receipt) int {
	points := 0

	//One point for every alphanumeric character in the retailer name
	for _, char := range receipt.Retailer {
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
		if len(trimmedDesc)%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd.
	day := strings.Split(receipt.PurchaseDate, "-")[2]
	dayInt, _ := strconv.Atoi(day)
	if dayInt%2 != 0 {
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

func findReceiptID(receipt receipt) int {
	for _, r := range receipts {
		if r.Retailer == receipt.Retailer && r.PurchaseDate == receipt.PurchaseDate && r.PurchaseTime == receipt.PurchaseTime && r.Total == receipt.Total && len(r.Items) == len(receipt.Items) {
			return r.ID
		}
	}

	return -1
}