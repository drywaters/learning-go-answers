package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int, 1)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go sendNumbers(ch)
	go sendNumbers(ch)
	go readNumbers(&wg, ch)

	wg.Wait()
}

func readNumbers(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 20; i++ {
		fmt.Println(i, "\t", <-ch)
	}
	wg.Done()
}

func sendNumbers(ch chan int) {
	for i := 0; i < 10; i++ {
		num := rand.Intn(1000)
		ch <- num
	}
}
