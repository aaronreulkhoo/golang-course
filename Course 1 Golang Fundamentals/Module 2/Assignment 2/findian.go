package main

import (
	"fmt"
	"strings"
)

func main() {
	var scanned string
	fmt.Printf("Please enter a string: ")
	_, err := fmt.Scanln(&scanned)
	if err != nil {
		fmt.Println("Error encountered")
	} else if strings.ToLower(string(scanned[0])) == "i" && strings.ToLower(string(scanned[len(scanned)-1])) == "n" && strings.ContainsAny(scanned, "aA") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
