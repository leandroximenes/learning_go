package lasagna

import "slices"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparation int) int {
	if averagePreparation < 1 {
		averagePreparation = 2
	}
	return len(layers) * averagePreparation
}

// TODO: define the 'Quantities()' function
func Quantities(ingredients []string) (int, float64) {
	var totalNoodles int
	var totalSauce float64
	for i := 0; i < len(ingredients); i++ {
		switch ingredients[i] {
		case "noodles":
			totalNoodles += 50
		case "sauce":
			totalSauce += 0.2
		}
	}

	return totalNoodles, float64(totalSauce)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	hasSecretIngredient := false
	for _, v := range friendsList {
		if !(slices.Contains(myList, v)) {
			hasSecretIngredient = true
		}
	}

	if hasSecretIngredient {
		myList = slices.Concat(myList, friendsList)
	}

	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountsNeeded []float64, qtde int) []float64 {
	panic("not implemented")
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
