package main

import (
	"fmt"
	"math"
)

const value = 20

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = math.MaxUint8
	smallI = math.MaxInt32
	bigI = math.MaxUint64

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Println("b: ", b)
	fmt.Println("smallI: ", smallI)
	fmt.Println("bigI: ", bigI)
}
