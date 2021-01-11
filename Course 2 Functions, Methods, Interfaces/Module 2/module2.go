package main

import (
	"fmt"
	"math"
)

/* M2.1.1 First Class Functions
Functions are a primary data type - can declare variables as functions, can be used as parameters/return values*/
func main() {
	funcVarTest()

	//returning functions
	o1 := makeDistOrigin(0, 0)
	o2 := makeDistOrigin(2, 2) //origin function with origin (2,2)
	fmt.Println(o1(2, 2))
	fmt.Println(o2(2, 2))

	//variadic functions
	fmt.Println(getMax(1, 2, 3, 4, 5))
	fmt.Println(getMax([]int{1, 2, 3, 4, 5}...)) // passing slice with ... suffix

	//defer usage
	deferTest()
}

// declares funcVar as a function which takes and return an integer
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

func funcVarTest() {
	funcVar = incFn
	fmt.Println(funcVar(1))
}

//applies function as argument
func applyFunc(funct func(int) int, val int) int {
	return funct(val)
}

//anonymous functions - nameless functions as arguments (similar to lambda functions)
func anonFunc() {
	_ = applyFunc(func(x int) int { return x + 1 }, 2)
}

//M2.1.2 Returning Functions
//For example, instead of introducing origin as another set of parameters
//Takes in the new origin and returns a new function calculating the distance between an origin point and another point
//Closure refers to the function + its environment
//When functions are passed/returned, their environment comes with them!!
func makeDistOrigin(oX, oY float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-oX, 2) + math.Pow(y-oY, 2))
	}
	return fn
}

//M2.1.3 Variadic and Deferred functions
//Variadic functions take in any number of arguments - treated as slice, hence can pass a slice with the ... suffix
func getMax(vals ...int) int {
	maxV := -1
	for _, val := range vals {
		if val > maxV {
			maxV = val
		}
	}
	return maxV
}

// Function calls can be deferred until surrounding function completes
// Arguments of a deferred call are evaluated immediately but the CALL is deferred
func deferTest() {
	i := 1
	defer fmt.Println(i + 1) // will print as 2, not 3
	i++
}
