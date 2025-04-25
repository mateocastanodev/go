package main

import (
	"fmt"
	"log"
)

func main() {
	var s string
	fmt.Printf("Write a string:")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	fmt.Println(s)
}
