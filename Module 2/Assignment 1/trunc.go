package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		scanned, err := scan()
		if err == nil {
			truncated := int(scanned)
			fmt.Printf("Your integer is %d.", truncated)
			break
		}
		fmt.Println("That's not a float!")
		stdin.ReadString('\n')

	}

}

func scan() (float64, error) {
	var scanned float64
	fmt.Printf("Please enter a float: ")
	_, err := fmt.Scan(&scanned)
	return scanned, err
}
