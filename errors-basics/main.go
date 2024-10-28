package main

import (
	"errors"
	"fmt"
)

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil
}

func div2(x, y int) int {
	return x / y
}

func div3(x, y int) int {
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}

func main() {
	result, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	i := 0
	result, _ = div(10, i)
	fmt.Println(result)

	result2 := div2(10, 2)
	fmt.Println(result2)

	result3 := div3(10, 0)
	fmt.Println(result3)
}
