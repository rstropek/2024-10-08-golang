package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(greeting string) {
	time.Sleep(3 * time.Second)
	fmt.Println(greeting)
}

func sayHelloWithWaitgroup(greeting string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Println(greeting)
}

func getValue() int {
	return 42
}

// NOT GOOD!
func getValueDelayed(result *int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	*result = 42
}

type Message struct {
	Name string
	Value int
}

func getValueWithChannel(name string, result chan<- Message) {
	time.Sleep(1 * time.Second)
	result <- Message{Name: name, Value: 42}
	result <- Message{Name: name, Value: 43}
}

func valueGenerator(result chan<- int, quit <-chan bool) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-quit:
			return
		default:
			result <- 42
		}
	}
}

func main() {
	/*
	wg := sync.WaitGroup{}
	wg.Add(2)
	go sayHelloWithWaitgroup("Hello, World!", &wg)
	go sayHelloWithWaitgroup("Hello, World!", &wg)
	wg.Wait()
	
	wg = sync.WaitGroup{}
	wg.Add(2)
	var a, b int
	go getValueDelayed(&a, &wg)
	go getValueDelayed(&b, &wg)
	wg.Wait()
	fmt.Printf("a: %d, b: %d\n", a, b)

	result := make(chan Message)
	go getValueWithChannel("a", result)
	go getValueWithChannel("b", result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	*/

	result := make(chan int)
	quit := make(chan bool)
	go func() {
		for x := range result {
			fmt.Println(x)
		}
	}()
	go valueGenerator(result, quit)

	time.Sleep(3 * time.Second)
	quit <- true
}
