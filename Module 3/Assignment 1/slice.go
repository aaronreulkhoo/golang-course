package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) //for clearing buffer
	var initialSize int = 3
	var currentSize int
	var slice = make([]int, initialSize)
	var scanned string

	for {
		fmt.Printf("Please enter an integer or 'X' to exit: ")
		fmt.Scan(&scanned)
		if strings.ToLower(scanned) == string('x') {
			fmt.Printf("Your final sorted slice is %v", slice)
			break
		} else {
			num, err := strconv.Atoi(scanned)
			if err == nil {
				if currentSize < initialSize { //find first 0 and replace with new number for first  entries
					for i, v := range slice {
						if v == 0 {
							slice[i] = num
							break
						}
					}
					currentSize++
				} else {
					slice = append(slice, num)
				}
				sort.Ints(slice)
				fmt.Printf("Your sorted slice is %v\n", slice)
			} else {
				fmt.Println("That's not an integer!")
			}
			stdin.ReadString('\n') //clear buffer
		}
	}
}
