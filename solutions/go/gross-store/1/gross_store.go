package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMap := map[string]int{}
    unitMap["quarter_of_a_dozen"] = 3
    unitMap["half_of_a_dozen"] = 6
    unitMap["dozen"] = 12
    unitMap["small_gross"] = 120
    unitMap["gross"] = 144
    unitMap["great_gross"] = 1728
    return unitMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
    if !exists {
        return false
    }
    bill[item] += value
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemInBill, isInBill := bill[item]
    unitValue, unitExists := units[unit]
    newValue := itemInBill - unitValue
    switch {
        case !isInBill || !unitExists || newValue < 0:
        	return false
        case newValue == 0:
        	delete(bill, item)
        	return true
        default:
        	bill[item] = newValue
        	return true
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
    if !exists {
        return 0, false
    }
    return value, true
}
