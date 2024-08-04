package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	emp1 := Employee{"Bob", "Builder", 1234}
	emp2 := Employee{
		firstName: "Ricky",
		lastName:  "Bobby",
		id:        999,
	}
	var emp3 Employee
	emp3.firstName = "Somebody"
	emp3.lastName = "Else"
	emp3.id = 0

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
