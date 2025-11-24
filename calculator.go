package main

import (
	"fmt"
)

func main() {
	fmt.Println("🧮 Simple Go Calculator")
	fmt.Println("Enter: number operator number")
	fmt.Println("Example: 5 + 3")
	fmt.Println("Operators: + - * /")
	fmt.Println("Type 'exit' to quit")
	fmt.Println()

	for {
		fmt.Print(" ")

		var num1 int
		var operator string
		var num2 int

		// Read input
		_, err := fmt.Scanln(&num1, &operator, &num2)
		if err != nil {
			fmt.Println("Error: Please use format: number operator number")
			continue
		}

		// Calculate
		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %d\n", result)
	}
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}
