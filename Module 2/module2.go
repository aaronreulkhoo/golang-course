package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Module 2")
	pointer1()
	pointer2()
	unicodeTest()
	stringsTest()
	strconvTest()
	controlFlowTest()
	scanTest()
}

// M2.1.1 Pointers - similar to C, & returns address of var,* returns data (deferencing)
func pointer1() {
	var x int = 1
	var y int
	var ip *int
	ip = &x //ip points to address of x
	y = *ip //y is now x
	fmt.Println(ip)
	fmt.Println(y)
}

// new() creates a variable and returns a pointer
func pointer2() {
	ptr := new(int)
	*ptr = 3
	fmt.Println(*ptr)
}

/* M2.1.2 Variable Scope - Blocks are sequences of declarations/statements within matching {}
Universe block refers to Go source code, package block refers to all source under main
Each block has environment of variables, builds a hierachy - similar to python */
var global string = "global"

func scope() { //this is a block, local variable cannot be referenced outside
	// var local string = "local"
}

/* M2.1.3 Deallocation - local variables in function calls (stack) are usually automatically collected,
global variables in heap must usually be collected manually like free() in c
M2.1.4 However, Go has automatic garbage collection despite being a compiled language!! :)*/

// M2.2.1 Comments, Printing - nothing special

// M2.2.2 Ints, Floats, Strings
// Ints and floats - casting similar to C, otherwise similar
var a int32 = 1
var b int16 = 2

// a=b will throw type error
func castingTest() {
	a = int32(b) // cast b into int32
}

// Strings - sequence of bytes (utf-8) that are read-only, each byte is a rune (basically a character)
var str string = "Hello"

//M2.2.3 Strings Packages
// unicode package provides set of functions to test runes e.g.:
const nihongo = "日本語"

func unicodeTest() {
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
		unicode.IsPunct(runeValue)
	}
}

// strings package provides functions to manipulate strings e.g.:
func stringsTest() {
	fmt.Println(strings.Contains(str, "Hel")) //true
}

//strconv package provides conversions to and from string represetnations of datas type
func strconvTest() {
	num := "32"
	fmt.Println(strconv.Atoi(num)) //converts ASCII string s to int
}

// M2.3.1 Constants - Expression whose value is known at compile time, inferred from value, cant be changed
const c1 = 1.3
const (
	c2 = 4
	c3 = "hi"
)

// iota - generate a set of related but distinct constants e.g. day and months
type Grades int

const (
	A Grades = iota //each constant is assigned a unique integer
	B
	C
	D
	F
)

//M2.3.2 Control Flow Structures
func controlFlowTest() {
	//If  - if <condition> {code}
	if true {
		fmt.Println("If statement true")
	} else {
		fmt.Println("Else statement")
	}
	//For Loop - for <init>; <terminatinv condition>; update {code}
	for i := 1; i < 2; i++ {
		fmt.Println(i)
	}
	//break breaks the current loop, continue skips the rest of the current iteration - normal

	c := 1
	//Switch - typical case, no need for breaks at end of case
	switch c {
	case 1:
		fmt.Println("1case1")
	case 2:
		fmt.Println("1case2")
	default:
		fmt.Println("1nocase")
	}

	//Tagless switch
	switch {
	case c > 1:
		fmt.Println("2case1")
	case c < -1:
		fmt.Println("2case2")
	default:
		fmt.Println("2nocase")
	}
}

//M2.2.3 Scan - Reads user input, takes pointer as argument, returns number of scanned items+pointer
func scanTest() {
	var scanned int
	fmt.Printf("Input a number:")
	num, err := fmt.Scan(&scanned)
	println(scanned, num, err)
}
