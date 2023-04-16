package main

import (
	"fmt"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("Hello, world!")
	return nil
}
