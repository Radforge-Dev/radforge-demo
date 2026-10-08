package shipping

// ShippingCost returns the shipping charge in cents.
// Orders of at least 5,000 cents qualify for free shipping.
func ShippingCost(subtotalCents int) int {
	if subtotalCents >= 5000 {
		return 0
	}
	return 599
}
