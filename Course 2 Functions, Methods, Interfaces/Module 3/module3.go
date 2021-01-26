package main

import (
	"fmt"
	"math"
)

func main() {
	assocTest()
	methodStructTest()
}

// M3.1.1 Classes and Encapsulation
// Universal idea - collection of data fields and functions that share well-defined responsibility

// M3.1.2 Method associations -  (Go has no official class keyword)
// However can associate methods with data types by using a receiver type defined before a function, use dot notation to call
type MyInt int

func (m MyInt) doubleInt() int {
	return int(m * 2)
}

func assocTest() {
	n := MyInt(3) // Note that a copy of n is being passed
	fmt.Println(n.doubleInt())
}

// M3.1.3 Structs with Methods
type Point struct {
	x float64
	y float64
}

func (p Point) distOrigin() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2))
}

func methodStructTest() {
	p1 := Point{3, 4}
	fmt.Println(p1.distOrigin())
}

// M3.2.1 Encapsulation - can define public functions to allow access to hidden data
// Start variables/functions to hide them in imports
/* Example to make 'constructor'
package data
type Point struct {
	x float64
	y float64
}

func(p* Point) initMe(x,y float64) {
	p.x=x
	p.y=y
}

// in main
package main
import "data"
func main() {
	var p data.Point
	p.initMe(3,4)
	p.otherMethodsInData
}
*/

// M3.2.1 Pointer Receivers
// Methods cannot change data inside the receiver
// Large receivers will copy everything (e.g. images which are huge arrays of ints)
// Use pointer receivers using obj *object
func (p *Point) initMe(x, y float64) {
	p.x = x
	p.y = y
}

//M3.2.3 Pointer Receivers, Rereferencing, Dereferencing
/* Notice that there is no need to dereference the object as it is automatic with dot notation
There is also no need to reference when calling the method
Good programming practice to standardise and either have all methods on a single type/struct
have pointer receivers (call by pointer) or normal (call by value - pass a copy)
*/
