package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	s := bufio.NewReader(os.Stdin)
	fmt.Printf("Write a string:")
	input, err := s.ReadString('\n')
	if err != nil {
		log.Printf("Error: %v", err)
	}
	lowerCase := strings.ToLower(input)
	if strings.Contains(lowerCase, "i") && strings.Contains(lowerCase, "a") && strings.Contains(lowerCase, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
