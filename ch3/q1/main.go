package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "こんにちは", "привет", "Mahalo"}
	firstTwo := greetings[:2]
	middleThree := greetings[1:4]
	lastTwo := greetings[3:]
	fmt.Println("firstTwo: ", firstTwo)
	fmt.Println("middleThree: ", middleThree)
	fmt.Println("lastTwo: ", lastTwo)
}
