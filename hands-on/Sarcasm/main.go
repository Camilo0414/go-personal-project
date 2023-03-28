// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// const message string = "I am good at coding"

// func main() {

// 	val := true
// 	var str string
// 	for _, i := range message {
// 		if unicode.IsLetter(i) {
// 			val = false
// 			if val != true {
// 				str += strings.ToLower(string(i))
// 			} else {
// 				strings.ToUpper(string(i))
// 			}
// 		}
// 		str += string(i)
// 	}

// 	fmt.Println(str)
// }

package main

import (
	"fmt"
	"unicode"
)

const message string = "I am good at coding"

func main() {

	isLower := true
	result := ""
	for _, c := range message {
		if unicode.IsLetter(c) {
			isLetter := true
			if unicode.IsUpper(c) {
				c += 'a' - 'A'
				isLetter = false
			}
			if isLower {
				if isLetter {
					result += string(c)
				} else {
					result += string(c - 'a' + 'A')
				}
			} else {
				if isLetter {
					result += string(c - 'a' + 'A')
				} else {
					result += string(c)
				}
			}
			isLower = !isLower
		} else {
			result += string(c)
		}
	}
	fmt.Println(result)
}
