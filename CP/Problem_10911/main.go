//Forming Quiz Teams

package main

import (
	"fmt"
	"math"
)

func main() {
	houses := [][]int{
		// N = len(houses)
		// (x,y) coords
		{1, 1},
		{8, 6},
		{6, 8},
		{1, 3},
	}
	fmt.Printf("The minimum cost is: %f\n", minCost(houses))
}

func minCost(xy [][]int) float64 {
	cost := 0.0

	if len(xy) > 1 && len(xy) <= 8 {
		myCoords := []int{}

		for i, x := range xy {
			for _, y := range x {
				myCoords := append(myCoords, y)
			}
			if i != 0 && i%2 == 0 {
				distance()
			}
		}
	} else {
		cost = 0.0
	}
	return cost
}

func distance(x, y, i, j float64) float64 {
	return math.Sqrt(math.Pow((i-x), 2) + math.Pow((j-y), 2))
}

func sum(ds []float64) float64 {
	sum_ds := 0.0
	for _, d := range ds {
		sum_ds += d
	}
	return sum_ds
}
