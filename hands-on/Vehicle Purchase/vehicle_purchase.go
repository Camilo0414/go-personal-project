package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return strings.Contains(kind, "car") || strings.Contains(kind, "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	v := ""
	if option1 < option2 {
		v = option1
	} else {
		v = option2
	}

	return v + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var value float64
	if age < 3 {
		value = originalPrice * 0.80
	} else if age >= 3 && age < 10 {
		value = originalPrice * 0.7
	} else {
		value = originalPrice * 0.5
	}
	return value
}
