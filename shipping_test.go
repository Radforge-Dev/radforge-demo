package shipping

import "testing"

func TestShippingCost(t *testing.T) {
	for _, tc := range []struct {
		name     string
		subtotal int
		want     int
	}{
		{"small order", 1000, 599},
		{"large order", 10000, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShippingCost(tc.subtotal); got != tc.want {
				t.Fatalf("ShippingCost(%d) = %d, want %d", tc.subtotal, got, tc.want)
			}
		})
	}
}
