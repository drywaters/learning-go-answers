package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	start := time.Now()
	people := make([]Person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{
			FirstName: "Bob",
			LastName:  "Builder",
			Age:       100,
		})
	}
	end := time.Now()
	fmt.Println(time.Duration(end.Sub(start)))
}
