package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var acceleration, velocity, displacement, time float64
	ls := []*float64{&acceleration, &velocity, &displacement, &time}
	vars := []string{"acceleration", "velocity", "displacement", "time"}
	i := 0
	stdin := bufio.NewReader(os.Stdin)
	for i < len(ls) {
		fmt.Printf("Please enter a value for %s: ", vars[i])
		scanned, _ := stdin.ReadString('\n')
		scanned = strings.Replace(scanned, "\r\n", "", -1) // comment out for UNIX systems
		num, err := strconv.ParseFloat(scanned, 64)
		if err == nil {
			*ls[i] = num
			i++
		} else {
			fmt.Println("That's not an integer!")
		}
	}
	DisplaceFn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Printf("The calculated displacement is %f", DisplaceFn(time))

}

// GenDisplaceFn : Generates a function which calculates displacement
func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
	DisplaceFn := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + velocity*time + displacement
	}
	return DisplaceFn
}
