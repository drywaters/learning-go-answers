package main

import (
	"errors"
	"fmt"
)

type ErrEmptyEmployeeField string

func (e ErrEmptyEmployeeField) Error() string {
	return string(e)
}

type Employee struct {
	FirstName, LastName string
	ID                  int
}

func (e Employee) validateEmployee() error {
	if e.FirstName == "" {
		return ErrEmptyEmployeeField("empty field: FirstName")
	} else if e.LastName == "" {
		return ErrEmptyEmployeeField("empty field: LastName")
	} else if e.ID == 0 {
		return ErrEmptyEmployeeField("empty field: ID")
	} else {
		return nil
	}
}

func main() {
	var e ErrEmptyEmployeeField
	e1 := Employee{"Bob", "Builder", 10}
	err := e1.validateEmployee()
	if err != nil && errors.As(err, &e) {
		fmt.Println("validation: ", err)
	}

	e2 := Employee{
		LastName: "Flay",
		ID:       1000,
	}
	err = e2.validateEmployee()
	if err != nil && errors.As(err, &e) {
		fmt.Println("validation: ", err)
	}

	e3 := Employee{
		FirstName: "Carl",
		ID:        1,
	}
	err = e3.validateEmployee()
	if err != nil && errors.As(err, &e) {
		fmt.Println("validation: ", err)
	}

	e4 := Employee{
		FirstName: "Carl",
		LastName:  "Sagan",
	}
	err = e4.validateEmployee()
	if err != nil && errors.As(err, &e) {
		fmt.Println("validation: ", err)
	}
}
