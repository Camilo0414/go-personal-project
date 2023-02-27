package main

import "fmt"

func main() {
	slice_ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, s := range slice_ints {
		if s%2 == 0 {
			fmt.Println(fmt.Sprint(s) + " is even")
		} else {
			fmt.Println(fmt.Sprint(s) + " is odd")
		}

	}
}
