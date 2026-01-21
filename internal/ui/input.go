package ui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// scanner is shared across all input operations to maintain consistent input handling
var scanner = bufio.NewScanner(os.Stdin)

// GetNumbers prompts for and validates two numeric inputs
// Returns both numbers and nil error on success, or zero values and error on failure
func GetNumbers() (float64, float64, error) {
	// Get first number
	fmt.Print("Enter first number: ")
	if !scanner.Scan() {
		return 0, 0, errors.New("Failed to read input")
	}
	firstInput := strings.TrimSpace(scanner.Text())
	num1, err := strconv.ParseFloat(firstInput, 64)
	if err != nil {
		return 0, 0, errors.New("Invalid input! Please enter a valid number.")
	}

	// Get second number
	fmt.Print("Enter second number: ")
	if !scanner.Scan() {
		return 0, 0, errors.New("Failed to read input")
	}
	secondInput := strings.TrimSpace(scanner.Text())
	num2, err := strconv.ParseFloat(secondInput, 64)
	if err != nil {
		return 0, 0, errors.New("Invalid input! Please enter a valid number.")
	}

	return num1, num2, nil
}

// GetMenuChoice prompts for and validates menu choice
// Returns choice (1-5) and nil error on success, or zero and error on failure
func GetMenuChoice() (int, error) {
	if !scanner.Scan() {
		return 0, errors.New("Failed to read input")
	}

	input := strings.TrimSpace(scanner.Text())
	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Invalid input! Please enter a number between 1-5.")
	}

	if choice < 1 || choice > 5 {
		return 0, errors.New("Invalid choice! Please select a number between 1-5.")
	}

	return choice, nil
}

// WaitForEnter waits for the user to press Enter
func WaitForEnter() {
	scanner.Scan()
}
