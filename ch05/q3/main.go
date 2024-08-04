package main

import "fmt"

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func prefixer(pre string) func(string) string {
	return func(s string) string {
		return pre + " " + s
	}
}
