package main

import "fmt"

func main() {
	var total int
	for i := 0; i < 10; i++ {
		// Change `total := to` `total =`
		total := total + i
		fmt.Println(total)
	}
	fmt.Println(total)
	// Did not print out the correct total because of shadowing
}
