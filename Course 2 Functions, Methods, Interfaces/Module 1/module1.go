package main

import (
	"encoding/json"
	"fmt"
)

/* M1.1 Functions
Main function has no return values/arguments */

func main() {
	fmt.Println(funcTest(4, 5))
	y := 1

	//M1.2
	callByVal(y)
	fmt.Printf("Call by Value - %d\n", y) // this remains as 1
	callByPointer(&y)
	fmt.Printf("Call by pointer - %d\n", y) //this is now 2

	//M1.3
	s1 := [1]int{1}
	passArray(&s1)
	s1b, _ := json.Marshal(s1) //quick way to convert arrays/slices to byte array, then stringify
	fmt.Println("s1 - " + string(s1b))
	s2 := []int{1}
	passSlice(s2)
	s2b, _ := json.Marshal(s2)
	fmt.Println("s2 - " + string(s2b))

}

// M1.1.2 Function parameters and return values - syntax
func funcTest(a, b int) int {
	return a * b
}

// M1.1.3 Calling by vlaue vs calling by reference
// Calling by value does not mutate variables outside of a function - makes a copy
// Matters for huge arrays or data structures
func callByVal(y int) {
	y = y + 1
}

func callByPointer(y *int) {
	*y = *y + 1
}

// M1.1.4 Passing arrays and slices
// Passing arrays is not recommended, pass slices instead - in general use slices in go
// slices contain a pointer to the array, passing a slice copies the pointer
func passArray(x *[1]int) {
	(*x)[0] = 2 //do something
}

func passSlice(s []int) {
	s[0] = 2 //do something
}

/* M1.2.1 Well written functions
Name functions well, not abstract but also not too verbose - try to
Strive to create one 'operation' with each function - don't merge functionality to increase developer convenience but decrease readability
Minimize parameters - Use structs to better define inputs by grouping up related data
Partition conditionals to reduce control flow complexity
*/
