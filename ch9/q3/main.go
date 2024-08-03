package main

import (
	"errors"
	"fmt"
)

type ErrEmptyEmployeeField struct {
	Errors []error
}

func (e ErrEmptyEmployeeField) Error() string {
	return errors.Join(e.Errors...).Error()
}

type Employee struct {
	FirstName, LastName string
	ID                  int
}

func (e Employee) validateEmployee() error {
	var errs []error
	if e.FirstName == "" {
		errs = append(errs, errors.New("FirstName empty"))
	}
	if e.LastName == "" {
		errs = append(errs, errors.New("LastName empty"))
	}
	if e.ID == 0 {
		errs = append(errs, errors.New("ID empty"))
	}
	if len(errs) > 0 {
		return ErrEmptyEmployeeField{errs}
	} else {
		return nil
	}
}

func main() {
	employees := []Employee{}
	var e ErrEmptyEmployeeField
	employees = append(employees, Employee{"Bob", "Builder", 10})
	employees = append(employees, Employee{
		LastName: "Flay",
		ID:       1000,
	})
	employees = append(employees, Employee{
		FirstName: "Carl",
		ID:        1,
	})
	employees = append(employees, Employee{
		FirstName: "Carl",
	})
	employees = append(employees, Employee{})
	for _, v := range employees {
		err := v.validateEmployee()
		if err != nil && errors.As(err, &e) {
			fmt.Printf("validation errors:\n%v\n\n", err)
		}
	}
}
