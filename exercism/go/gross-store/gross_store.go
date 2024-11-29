package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	gross := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return gross
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	qtde, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += qtde
	return true

}

// RemoveItem removes an item from customer bill.
// RemoveItem(bill, units, "carrot", "dozen")
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	qtdeInBill, existsInBill := bill[item]
	untQtde, existsInUnits := units[unit]
	NewqtdeInBill := qtdeInBill - untQtde

	if !existsInBill || !existsInUnits || NewqtdeInBill < 0 {
		return false
	}

	if NewqtdeInBill == 0 {
		delete(bill, item)
	} else {
		bill[item] = NewqtdeInBill
	}

	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qtde, exists := bill[item]
	return qtde, exists
}
