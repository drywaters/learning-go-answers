package main

import "fmt"

type DoubleNums interface {
	int | float64 | float32
}

func main() {
	i := 10
	f := 5.55
	fmt.Println(Double(i))
	fmt.Println(Double(f))
}

func Double[T DoubleNums](i T) T {
	return i * i
}
