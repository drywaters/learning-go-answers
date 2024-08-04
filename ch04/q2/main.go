package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}

	for _, n := range nums {
		if n%2 == 0 && n%3 == 0 {
			fmt.Println("Six!")
		} else if n%2 == 0 {
			fmt.Println("Two!")
		} else if n%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never mind")
		}
	}
}
