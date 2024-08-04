package main

import "fmt"

func main() {
	strs := []string{"one", "two", "three"}

	fmt.Println("Strs before update: ", strs)
	UpdateSlice(strs, "four")
	fmt.Println("Strs after update: ", strs)

	fmt.Println("Strs before grow: ", strs)
	GrowSlice(strs, "five")
	fmt.Println("Strs after grow: ", strs)
}

func UpdateSlice(strs []string, s string) {
	strs[len(strs)-1] = s
}

func GrowSlice(strs []string, s string) {
	strs = append(strs, s)
}
