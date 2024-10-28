package main

import "fmt"

type DivisionError struct {
	message string
}

func (e *DivisionError) Error() string {
	return e.message
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, &DivisionError{message: "division by zero"}
	}

	return x / y, nil
}

func main() {
	result, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
