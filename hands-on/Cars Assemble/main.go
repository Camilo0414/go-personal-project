package main

func main() {
	CalculateCost(37)

}

func CalculateCost(carsCount int) uint {

	return uint((carsCount%10)*10000 + (carsCount/10)*95000)
}
