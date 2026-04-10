package thefarm

import (
	"errors"
	"fmt"
)

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

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cowsNumber int) error {
	if cowsNumber > 0 {
		return nil
	}

	var customMessage string
	if cowsNumber < 0 {
		customMessage = "there are no negative cows"
	}

	if cowsNumber == 0 {
		customMessage = "no cows don't need food"
	}

	erroMessage := fmt.Sprintf("%d cows are invalid: %s", cowsNumber, customMessage)

	return errors.New(erroMessage)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
