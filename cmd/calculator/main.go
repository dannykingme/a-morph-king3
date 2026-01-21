package main

import (
	"fmt"

	"github.com/your-org/a-morph-king3-go/internal/calculator"
	"github.com/your-org/a-morph-king3-go/internal/ui"
)

func main() {
	fmt.Println("Welcome to the Simple Calculator!")

	for {
		// Display menu
		ui.DisplayMenu()

		// Get menu choice
		choice, err := ui.GetMenuChoice()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Check for exit
		if choice == 5 {
			fmt.Println("Thank you for using the calculator! Goodbye!")
			break
		}

		// Get numbers for calculation
		num1, num2, err := ui.GetNumbers()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// Perform operation based on choice
		var result float64
		switch choice {
		case 1:
			result = calculator.Add(num1, num2)
			fmt.Printf("%g + %g = %g\n", num1, num2, result)
		case 2:
			result = calculator.Subtract(num1, num2)
			fmt.Printf("%g - %g = %g\n", num1, num2, result)
		case 3:
			result = calculator.Multiply(num1, num2)
			fmt.Printf("%g * %g = %g\n", num1, num2, result)
		case 4:
			result, err = calculator.Divide(num1, num2)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("%g / %g = %g\n", num1, num2, result)
			}
		}

		// Press Enter to continue
		fmt.Print("\nPress Enter to continue...")
		ui.WaitForEnter()
	}
}
