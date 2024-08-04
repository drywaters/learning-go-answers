package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {
	p1 := MakePerson("Bob", "Builder", 100)
	p2 := MakePerson("Other", "Person", 10)
	fmt.Println(p1)
	fmt.Println(p2)
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}
