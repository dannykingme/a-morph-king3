package ui

import "fmt"

// DisplayMenu prints the calculator menu to stdout
func DisplayMenu() {
	fmt.Println("\n========== Simple Calculator ==========")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Println("5. Exit")
	fmt.Println("=======================================")
	fmt.Print("Enter your choice (1-5): ")
}
