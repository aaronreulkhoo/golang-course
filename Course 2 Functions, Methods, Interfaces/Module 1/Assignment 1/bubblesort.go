package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sli []int
	var maxSize int = 10
	stdin := bufio.NewReader(os.Stdin)
	for len(sli) < maxSize {
		fmt.Printf("Please enter integer %d or leave it empty to finish the sequence: ", len(sli)+1)
		scanned, _ := stdin.ReadString('\n')
		scanned = strings.Replace(scanned, "\r\n", "", -1) // comment out for UNIX systems
		if scanned == "" {
			break
		}
		num, err := strconv.Atoi(scanned)
		if err == nil {
			sli = append(sli, num)
		} else {
			fmt.Println("That's not an integer!")
		}
	}
	slib, _ := json.Marshal(sli)
	fmt.Println("The array provided is " + string(slib))
	fmt.Println("Commencing Bubble Sort...")
	BubbleSort(sli)
	slib, _ = json.Marshal(sli)
	fmt.Println("The final sorted array is " + string(slib))
}

// BubbleSort : Implements the bubble sort algorithm to sort an input slice
func BubbleSort(sli []int) {
	sortedFlag := true
	for sortedFlag {
		sortedFlag = false
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				sortedFlag = true
			}
		}
	}
}

// Swap : swaps the ith and i+1th element of an input slice
func Swap(sli []int, index int) {
	sli[index+1], sli[index] = sli[index], sli[index+1]
}
