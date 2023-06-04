package models

type item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}