package main

import (
	"fmt"
	"math"
	"sync"
)

var rootCacheFn = sync.OnceValue(sqrts)

func main() {
	for i := 1000; i < 100_000; i += 1000 {
		fmt.Println(i, rootCacheFn()[i])
	}
}

func sqrts() map[int]float64 {
	roots := map[int]float64{}
	for i := 0; i < 100_000; i++ {
		roots[i] = math.Sqrt(float64(i))
	}
	return roots
}
