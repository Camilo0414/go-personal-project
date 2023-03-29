package techpalace

import (
	"strings"
	"unicode"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	st := strings.Repeat("*", numStarsPerLine) + "\n" +
		welcomeMsg + "\n" +
		strings.Repeat("*", numStarsPerLine)

	return st
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.Split(oldMsg, "\n")
	msg := strings.TrimFunc(newMsg[1], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '%'
	})
	return msg
}
