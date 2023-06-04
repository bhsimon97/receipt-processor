package main

type item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

type allReceipts []receipt

type pointsResponse struct {
	Points int `json:"points"`
}

type processReceiptResponse struct {
	ID int `json:"id"`
}

type receipt struct {
	ID           int     `json:"id,string"`
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []item  `json:"items,string"`
	Total        float64 `json:"total,string"`
}