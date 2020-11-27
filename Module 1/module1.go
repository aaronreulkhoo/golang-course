package main

// main package starts hierachy

// Import Declarations
import "fmt"

// Variable Declaration - keyword name type
var x int
var y string
var z float32

// Type Declarations - aliasing datatypes
type Celsius float64

var temp Celsius

// Functions - func function_name() return_type
func main() {
	// Short Variable Declaration + Initialization - infers type, can only do inside function
	instant := "Hello, world!"
	fmt.Println(instant)
}
