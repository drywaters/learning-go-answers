package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(i, j int) (int, error) { return i + j, nil }
func sub(i, j int) (int, error) { return i - j, nil }
func mul(i, j int) (int, error) { return i * j, nil }
func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println("invalid expression:", exp)
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := exp[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}

		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println("invalid operation:", err)
			continue
		}
		fmt.Println(result)
	}

}
