package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	interfaceTest()
	nilValueTest()
	implementTest()
	emptyInterfaceTest([]int{1, 2, 3, 4, 5})
	typeAssertionTest()
	errorTest()
}

// M4.1.1 Polymorphism - ability for an object to take different forms depending on context
// As introduced, Golang does not have inheritance like in other OO languages
// M4.1.2 Interfaces - set of method signatures - name, parameters, return values, doesnt actually implement methods, no data

type shape2d interface {
	doSomething() float64
}

type triangle struct {
	name string
}

// A type satisfies an interface if it defines all methods in the interface, no need to state explicitly like in other languages
func (t triangle) doSomething() {
	fmt.Println(t.name)
}

// M4.1.3 Concrete vs Interface types - Interfaces are linked to concrete types, defined by dynamic type and dynamic value
// eg. Dog is a concrete type which satisfies the Speaker interface
// Dynamic type is Dog, Dynamic value is d1 variable
type animal interface {
	bark()
	owner()
}
type dog struct {
	name string
}

func (d dog) bark() {
	fmt.Println("woof")
}
func interfaceTest() {
	var s1 animal
	d1 := dog{"Brian"}
	s1 = &d1
	s1.bark()
}

// Interfaces can have nil dynamic value - can still call methods, however need to define behaviour

func (d *dog) owner() {
	if d == nil {
		fmt.Println("I'm a nameless dog")
	} else {
		fmt.Printf("Owned dog with name %s\n", d.name)
	}
}
func nilValueTest() {
	var s1 animal
	var d1 *dog
	s1 = d1
	s1.owner()
}

// M4.2.1 Using interfaces - express similarity between type - allows function to take in multiple types

type geometry interface {
	area() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
}
func implementTest() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

// Empty interfaces specify no method - all types accepted - use for universal functions
func emptyInterfaceTest(val interface{}) {
	fmt.Println(val)
}

// M4.2.2 Type Assertions - Interfaces conceal differences, however sometimes need to extract underlying concrete type (disambiguation)
func checkShape(g geometry) {
	switch sh := g.(type) {
	case rect:
		fmt.Printf("I am a rectangle %v\n", sh)
	case circle:
		fmt.Printf("I am a circle %v\n", sh)
	}
}

func typeAssertionTest() {
	r := rect{width: 3, height: 4}
	checkShape(r)
}

// M4.2.3 - Error interface - common interface returned as second argument
// fmt package calls error() method to generate string to print
/* type error interface{
	Error() string
} */
func errorTest() {
	_, err := os.Open("none.txt")
	if err != nil {
		fmt.Println(err)
	}
}
