package main

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