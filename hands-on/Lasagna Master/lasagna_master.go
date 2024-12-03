package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(slice []string, avg int) int {
	var time int

	if avg == 0 {
		time = len(slice) * 2
	} else {
		time = len(slice) * avg
	}

	return time
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (int, float64) {
	// grams := 50
	// lts := 0.2

	// qntNoodles := 0
	// qntSauce := 0
	return 0, 0.0
}

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
