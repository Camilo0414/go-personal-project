package main

import (
	"fmt"
	"sort"
)

const message string = "ABBCFFEB"

func main() {
	input := []byte(message)
	sort.Slice(input, func(i, j int) bool {
		return input[j] > input[i]
	})

	i, j := int(input[0]), int(input[len(input)-1])

	complete_range := make([]byte, j-i+1)
	for k := range complete_range {
		complete_range[k] = byte(i + k)
	}
	var val byte
	for a := range complete_range {
		for b := range input {
			if complete_range[a] == input[b] {
				val = 0
			}
		}
	}
	fmt.Println(complete_range, val)
}
