// Package weather provides forecast for a city (?).
package weather

// CurrentCondition is a nice var.
var CurrentCondition string

// CurrentLocation is a nicest var :).
var CurrentLocation string

// Forecast makes many things.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
