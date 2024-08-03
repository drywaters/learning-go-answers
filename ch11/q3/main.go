package main

import (
	_ "embed"
	"flag"
	"fmt"
)

//go:embed english_rights.txt
var engRights string

//go:embed mas_rights.txt
var masRights string

func main() {
	var lang string
	flag.StringVar(&lang, "language", "english", "which language: english/mas")
	flag.Parse()

	if lang == "english" {
		fmt.Println(engRights)
	} else if lang == "mas" {
		fmt.Println(masRights)
	} else {
		fmt.Println("Language not understood:  Select 'english' or 'mas'")
	}
}
