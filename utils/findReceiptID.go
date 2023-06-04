package utils

func findReceiptID(receipt receipt) int {
	for _, r := range receipts {
		if r.Retailer == receipt.Retailer && r.PurchaseDate == receipt.PurchaseDate && r.PurchaseTime == receipt.PurchaseTime && r.Total == receipt.Total && len(r.Items) == len(receipt.Items) {
			return r.ID
		}
	}

	return -1
}