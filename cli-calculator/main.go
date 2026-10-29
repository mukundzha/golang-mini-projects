
package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	for {
		fmt.Println("\nWelcome to Calculator!")
		fmt.Println("Enter first number (or 0 to exit):")
		fmt.Scan(&num1)

		if num1 == 0 {
			fmt.Println("Goodbye!")
			break
		}

		fmt.Println("Enter operator (+, -, *, /):")
		fmt.Scan(&operator)

		fmt.Println("Enter second number:")
		fmt.Scan(&num2)

		if operator == "+" {
			fmt.Println("Result:", num1+num2)

		} else if operator == "-" {
			fmt.Println("Result:", num1-num2)

		} else if operator == "*" {
			fmt.Println("Result:", num1*num2)

		} else if operator == "/" {
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero!")
			} else {
				fmt.Println("Result:", num1/num2)
			}

		} else {
			fmt.Println("Invalid operator!")
		}
	}
}

