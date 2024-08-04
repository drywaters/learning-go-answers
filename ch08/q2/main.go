package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (p PrintInt) String() string {
	return strconv.Itoa(int(p))
}

type PrintFloat float64

func (p PrintFloat) String() string {
	return strconv.FormatFloat(float64(p), byte('f'), 2, 64)
}

func main() {
	var i PrintInt = 10
	var f PrintFloat = 1000.55555
	PrintNumber(i)
	PrintNumber(f)
}

func PrintNumber[T Printable](p T) {
	fmt.Println(p)
}
