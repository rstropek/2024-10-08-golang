package main

import "fmt"

func setTo42(px *int) {
	*px = 42
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	x := 42
	px := &x
	*px = 43
	fmt.Printf("x: %v, px: %v\n", x, px)

	*px *= 2
	fmt.Println(x)

	px = new(int)
	setTo42(px)
	fmt.Println(*px)

	pperson := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	pperson.Age += 1
	fmt.Println(*pperson)
}
