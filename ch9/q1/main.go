package main

import (
	"fmt"
)

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrInvalidID = Sentinel("invalid id")
)

func main() {

	badId := -10
	goodID := 100

	err := checkID(badId)
	if err != nil {
		if err == ErrInvalidID {
			fmt.Println("id must be positive number: ", badId)
		}
	} else {
		fmt.Println("Good id found", goodID)
	}

	err = checkID(goodID)
	if err != nil {
		if err == ErrInvalidID {
			fmt.Println("id must be positive number: ", badId)
		}
	} else {
		fmt.Println("Good id found: ", goodID)
	}
}

func checkID(id int) error {
	if id < 0 {
		return ErrInvalidID
	} else {
		return nil
	}
}
