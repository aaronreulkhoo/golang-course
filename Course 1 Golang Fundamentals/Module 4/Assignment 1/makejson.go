package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := make(map[string]string)
	// get inputs
	fmt.Printf("Please enter a name: ")
	name, _ := stdin.ReadString('\n')
	name = strings.Replace(name, "\r\n", "", -1)
	fmt.Printf("Please enter an address: ")
	address, _ := stdin.ReadString('\n')
	address = strings.Replace(address, "\r\n", "", -1)
	m["name"], m["address"] = name, address
	// marshal
	byteArray, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
