package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), time.Second*time.Duration(2), errors.New("timeout reached"))
	defer cancel()
	for i := 0; ; i++ {
		if err := context.Cause(ctx); err != nil {
			fmt.Println(err)
			break
		}
		n1 := rand.Intn(100_000_000)
		n2 := rand.Intn(100_000_000)
		sum := n1 + n2
		fmt.Printf("Iteration: %d\t sum: %d\n", i, sum)
		if sum == 1234 {
			fmt.Println("Found sum of 1234")
			break
		}
	}
}
