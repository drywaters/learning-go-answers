package main

import "fmt"

const value = 20

func main() {
	var i int
	var f float64

	i = value
	f = value

	fmt.Printf("i: %d\n", i)
	fmt.Printf("f: %.2f\n", f)
}
