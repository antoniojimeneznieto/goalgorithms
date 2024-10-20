package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe()
	sum(10, 20)

	var result, remainder, err = division(10, 3)
	switch {
	case err != nil:
		fmt.Printf("Error: %s\n", err.Error())
	case remainder == 0:
		fmt.Printf("Division result: %d\n", result)
	default:
		fmt.Printf("Division result: %d, remainder: %d\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Remainder is zero")
	case 1, 2:
		fmt.Println("Remainder is either 1 or 2")
	default:
		fmt.Println("Remainder is neither 0, 1, nor 2")
	}

}

func printMe() {
	fmt.Println("Hello, World!")
}

func sum(a int, b int) int {

	result := a + b
	fmt.Println("Sum of", a, "and", b, "is", result)
	return result
}

func division(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Division by zero is not allowed")
		return 0, 0, err
	}
	result := a / b
	remainder := a % b

	return result, remainder, err
}
