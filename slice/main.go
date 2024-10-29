package main

import "fmt"

func main() {
	// Array
	const length = 5
	arr := [length]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(arr[0])
	fmt.Println(arr[1:])
	fmt.Println(arr[:1])
	// arr = append(arr, 6) // Does not work
	printAllElements(arr[:])

	// Slice
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2
	slice = []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(slice[0])
	fmt.Println(slice[1:])
	fmt.Println(slice[:1])
	printAllElements(slice)

	fmt.Println("===")
	myOtherSlice := slice[1:3]
	slice[2] = 10
	printAllElements(myOtherSlice)
}

func printAllElements(arr []int) {
	for _, element := range arr {
		fmt.Println(element)
	}
}
