package main

import "fmt"

func main() {
	var message string
	fmt.Scan(&message)
	for i := int(minRune(message)); i < int(maxRune(message)); i++ {
		if !containsRune(message, i) {
			fmt.Print(string(i))
		}
	}
}

func minRune(s string) rune {
	min := rune(s[0])
	for _, c := range s {
		if c < min {
			min = c
		}
	}
	return min
}

func maxRune(s string) rune {
	max := rune(s[0])
	for _, c := range s {
		if c > max {
			max = c
		}
	}
	return max
}

func containsRune(s string, r int) bool {
	for _, c := range s {
		if int(c) == r {
			return true
		}
	}
	return false
}
