package main

import (
	"errors"
	"fmt"
	"strconv"
)

func div2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func div(a, b int) int {
	return a / b
}


func repl() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	var a, b string
	for {
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		aNum, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("You entered an invalid number for the first number:", err)
			continue
		}

		bNum, err := strconv.Atoi(b)
		if err != nil {
			fmt.Println("You entered an invalid number for the second number:", err)
			continue
		}

		result := div(aNum, bNum)
		fmt.Printf("%d / %d = %d\n\n", aNum, bNum, result)
	}
}

func main() {
	for {
		repl()
	}
}
