package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(numbersChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	numbersChannel <- rand.Intn(100)
}

func main() {
	numbersChannel := make(chan int)
	/*
	numberOfProducers := 5 // In real life, this number would be dynamic
	wg := sync.WaitGroup{}
	wg.Add(numberOfProducers)

	for i := 0; i < numberOfProducers; i++ {
		go produce(numbersChannel, &wg)
	}

	doneChannel := make(chan bool)
	go func() {
		for elem := range numbersChannel {
			fmt.Println(elem)
		}

		// Not so good, for demonstration purposes only
		// for {
		// 	value, open := <-numbersChannel
		// 	if !open {
		// 		break
		// 	}
		// 	fmt.Println(value)
		// }

		doneChannel <- true
	}()

	wg.Wait()
	close(numbersChannel)
	<-doneChannel
	*/

	wg := sync.WaitGroup{}
	wg.Add(1)
	go produce(numbersChannel, &wg)
	select {
	case m := <-numbersChannel:
		fmt.Println(m)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}
