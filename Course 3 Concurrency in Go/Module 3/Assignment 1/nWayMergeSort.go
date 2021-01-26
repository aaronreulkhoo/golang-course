package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := []int{}
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Please enter an integer, or leave it empty to finish the array: ")
		scanned, _ := stdin.ReadString('\n')
		scanned = strings.Replace(scanned, "\r", "", -1)
		scanned = strings.Replace(scanned, "\n", "", -1)
		if scanned == "" {
			break
		}
		num, err := strconv.Atoi(scanned)
		if err == nil {
			arr = append(arr, num)
		} else {
			fmt.Println("That's not an integer!")
		}
	}

	workers := 4
	channels := make([]chan []int, workers)
	prevInd, nextInd := 0, 0
	fmt.Printf("There are currently %d goroutines with arrays:\n", workers)
	for i := 0; i < workers; i++ {
		prevInd = nextInd
		nextInd += len(arr) / workers
		if len(arr)%workers-i > 0 {
			nextInd++
		}
		channels[i] = make(chan []int)
		defer close(channels[i])
		go subSort(arr[prevInd:nextInd], channels[i])
	}

	subArrays := make([][]int, workers)
	for i := 0; i < workers; i++ {
		subArrays[i] = <-channels[i]
	}
	fmt.Printf("Performing %d-way merge...\n", workers)
	// N-way merge
	result := []int{}
	indexes := make([]int, workers)
	ints := make([]int, workers)
	for l := 0; l < len(arr); l++ {
		for s := 0; s < workers; s++ {
			if indexes[s] < len(subArrays[s]) {
				ints[s] = subArrays[s][indexes[s]]
			} else {
				ints[s] = math.MaxInt64
			}
		}
		minVal := math.MaxInt64
		minDex := 0
		for i, v := range ints {
			if v < minVal {
				minVal = v
				minDex = i
			}
		}
		indexes[minDex]++
		result = append(result, minVal)

	}
	fmt.Println("The sorted array is", result)

}

func subSort(subArray []int, ch chan []int) {
	sort.Ints(subArray)
	fmt.Println(subArray)
	ch <- subArray
}
