package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)




type eatOrKeep interface {
	shouldEat() bool
}

type lentil struct {
	isGood bool
}

func (l lentil) shouldEat() bool {
	return !l.isGood
}

type bird struct{}

func (p bird) process(item []eatOrKeep) []eatOrKeep {
	result := []eatOrKeep{}
	for _, lentil := range item {
		if !lentil.shouldEat() {
			result = append(result, lentil)
		}
	}
	return result
}

func (p bird) processStrategy(item []eatOrKeep, eatRule func(eatOrKeep) bool) []eatOrKeep {
	result := []eatOrKeep{}
	for _, lentil := range item {
		if !eatRule(lentil) {
			result = append(result, lentil)
		}
	}
	return result
}

func filter[T any](items []T, filterRule func(i T) bool) []T {
	result := []T{}
	for _, item := range items {
		if !filterRule(item) {
			result = append(result, item)
		}
	}
	return result
}

type myBagItem[T constraints.Integer] struct {
	item T
	next *myBagItem[T]
}

type apple struct {
	rotten bool
}

func (a apple) shouldEat() bool {
	return a.rotten
}

type rabbit struct{}

func (r rabbit) process(apples []apple) []apple {
	result := []apple{}
	for _, apple := range apples {
		if !apple.shouldEat() {
			result = append(result, apple)
		}
	}
	return result
}

func filterChannel[T any](items <-chan T, filterRule func(i T) bool) <-chan T {
	result := make(chan T)
	go func() {
		for item := range items {
			if filterRule(item) {
				result <- item
			}
		}
	}()
	return result
}

type myInterface[T any] interface {
	getItem() T
}

func main() {
	lentils := []eatOrKeep{
		lentil{isGood: true},
		lentil{isGood: false},
	}

	bird := bird{}
	result := bird.processStrategy(lentils, func(item eatOrKeep) bool {
		return item.shouldEat()
	})

	fmt.Println(result)

	numbers := []int{1, 2, 3, 4, 5}
	// Filter even numbers
	evenNumbers := filter[int](numbers, func(item int) bool {
		return item%2 == 0
	})
	fmt.Println(evenNumbers)

	bag := myBagItem[int]{
		item: 1,
		next: nil,
	}
	if bag.item%2 == 0 {
		fmt.Println("Even number")
	}
}
