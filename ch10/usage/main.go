package main

import (
	"fmt"

	"github.com/drywaters/adder/v2"
)

func main() {
	a := 10.0
	b := 20.5

	fmt.Println(adder.Add(a, b))
}
