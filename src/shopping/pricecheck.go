package shopping

import "shopping/db"

func PriceCheck(itemID int) (float64, bool) {
	item := db.LoadItem(itemID)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
