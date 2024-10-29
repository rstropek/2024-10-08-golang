package main

import (
	"fmt"
	"iter"
)

func main() {
	simpleIterator()
	fmt.Println("---")
	endlessIterator()
}

func simpleIterator() {
	// Seq[int]
	//var iterator iterator.Seq[int]
	iterator := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				break
			}
		}
	}

	for v := range Map(
		Filter(iterator, func(v int) bool { return v%2 == 0 }),
		func(v int) int { return v * v }) {
		fmt.Println(v)
		if v == 6 {
			break
		}
	}

	next, stop := iter.Pull(iterator)
	for v, ok := next(); ok; v, ok = next() {
		fmt.Println(v)
		if v == 6 {
			stop()
			break
		}
	}
}

func pull_iterator() {
	pull_iter := func() (func() (int, bool), func()) {
		current := 0

		next := func() (int, bool) {
			if current >= 10 {
				return 0, false
			}
			current++
			return current, true
		}

		stop := func() {
			current = 0
		}

		return next, stop
	}

	next, stop := pull_iter()
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		fmt.Println(v)
	}
}

func Filter[V any](iter iter.Seq[V], fn func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range iter {
			if fn(v) && !yield(v) {
				break
			}
		}
	}
}

func Map[V, W any](iter iter.Seq[V], fn func(V) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for v := range iter {
			if !yield(fn(v)) {
				break
			}
		}
	}
}

func endlessIterator() {
	iter := func(yield func(int) bool) {
		i := 0
		for {
			if !yield(i) {
				break
			}
			i++
		}
	}

	skip := 10
	take := 5
	for v := range iter {
		if v < skip {
			continue
		}
		fmt.Println(v)
		if v >= skip+take-1 {
			break
		}
	}
}
