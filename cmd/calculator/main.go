package main

import (
	"fmt"

	"github.com/your-org/a-morph-king3-go/internal/calculator"
)

func main() {
	// Minimal main.go to validate architecture
	// This demonstrates that operations can be called and error handling works

	fmt.Println("Testing calculator operations...")

	// Test Add
	sum := calculator.Add(5.0, 3.0)
	fmt.Printf("5 + 3 = %g\n", sum)

	// Test Subtract
	diff := calculator.Subtract(10.0, 4.0)
	fmt.Printf("10 - 4 = %g\n", diff)

	// Test Multiply
	product := calculator.Multiply(6.0, 7.0)
	fmt.Printf("6 * 7 = %g\n", product)

	// Test Divide (success case)
	quotient, err := calculator.Divide(20.0, 4.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("20 / 4 = %g\n", quotient)
	}

	// Test Divide (error case - division by zero)
	_, err = calculator.Divide(10.0, 0.0)
	if err != nil {
		fmt.Printf("Expected error for division by zero: %v\n", err)
	}

	fmt.Println("\nArchitecture validation complete!")
}
