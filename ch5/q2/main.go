package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "file", "", "filename to get length for")
	flag.Parse()
	length, err := fileLen(name)
	if err != nil {
		fmt.Println("unable to read file:", err)
		os.Exit(1)
	}
	fmt.Println("File length is: ", length)
}

func fileLen(name string) (int, error) {
	file, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	bytes := make([]byte, 1000)
	var length int
	for {
		n, err := reader.Read(bytes)
		length = length + n
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			if err == io.EOF {
				return length, nil
			}
		}
	}
}
