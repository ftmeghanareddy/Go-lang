package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operator string

	fmt.Print("Enter first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	switch operator {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
