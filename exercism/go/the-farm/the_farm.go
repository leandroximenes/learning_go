package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, cowsNumber int) (float64, error) {
	fodderAmount, err := fc.FodderAmount(cowsNumber)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	var divideFood float64 = fodderAmount * fatteningFactor / float64(cowsNumber)
	return divideFood, nil

}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, cowsNumber int) (float64, error) {
	if cowsNumber > 0 {
		return DivideFood(fc, cowsNumber)
	} else {
		return 0, errors.New("invalid number of cows")
	}

}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows returns nil if the cow count is valid, or *InvalidCowsError otherwise.
func ValidateNumberOfCows(cowsNumber int) error {
	if cowsNumber > 0 {
		return nil
	}

	// Create a default variable
	var erroResponse *InvalidCowsError = &InvalidCowsError{
		cows: cowsNumber, message: "no cows don't need food",
	}

	if cowsNumber < 0 {
		erroResponse.message = "there are no negative cows"
	}

	return erroResponse
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
