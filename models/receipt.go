package models

type receipt struct {
	ID           int     `json:"id,string"`
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []item  `json:"items,string"`
	Total        float64 `json:"total,string"`
}