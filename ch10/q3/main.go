// Package adder adds together two integers
package adder

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds together two [Number] returning their sum
//
// More information on adding can be found at [Mathisfun].
// [Mathisfun]: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
