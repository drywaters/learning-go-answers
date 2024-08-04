package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go sendNumbers(ch1)
	go sendNumbers(ch2)

loop:
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("channel 1", "\t", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("channel 2", "\t", v)
		default:
			if ch1 == nil && ch2 == nil {
				break loop
			}
		}
	}
}

func sendNumbers(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		num := rand.Intn(1000)
		ch <- num
	}
}
