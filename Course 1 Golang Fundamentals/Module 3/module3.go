package main

import "fmt"

func main() {
	arraysTest()
	slicesTest()
	varSlicesTest()
	mapsTest()
	structsTest()
}

// Composite Data Types
// M3.1.1 Arrays -[number] comes before type

func arraysTest() {
	// Initialization
	var x [5]int = [5]int{1, 2, 3, 4, 5} //full assignment
	z := [...]int{1, 2, 3}               //array literal, length inferred

	z[0] = 2          //asignment
	fmt.Println(x[1]) // arrays initialized to 1

	for index, value := range x { //python enumerate
		_, _ = index, value //dosomething
	}
}

// M3.1.2 Slices - a fragment of an existing array, size is variable - similar to python slicing
// Has 3 properties: Pointer -> start of slice ; Length -> #elements ; Capacity -> maximum possible #elements (array length - position)
func slicesTest() {
	arr := [...]string{"a", "b", "c"}
	s1 := arr[1:3] // holds arr[1], arr[2]
	fmt.Printf("length, capacity of slice = %d,%d\n", len(s1), cap(s1))
	s1[0] = "new" //slices change the underlying array
	fmt.Println(arr[1])

	s2 := []int{1, 2, 3} //slice literal, creates underlying array
	fmt.Println(s2)
}

//M3.1.3 Variable Slices - using make(type,length,capacity)
func varSlicesTest() {
	var s1 = make([]string, 10)    //creates slice+array of length 10
	var s2 = make([]string, 5, 15) //creates slice of length 10, capacity (length of array) of 15
	s1 = append(s1, "new")         //append(slice,element) adds elements to slice and expand underlying array if necessary
	fmt.Println(s1)
	fmt.Println(s2)
}

//M3.2.1 Hash Tables - each value associated with unique key mapped by hash function
//M3.2.2 Maps - Golang's implementation of hash tables - var name map[keytype][valtype]
func mapsTest() {
	m1 := make(map[string]string)   //using var m1 map[string]string will not initialise the hash - cannot assign!
	m2 := map[string]int{"test": 0} //map literal
	m1["hello"] = "world"           //assignment
	delete(m1, "hello")             //deletion
	v, p := m2["none"]              //double assignment - v is value, p is a boolean indicating presence of key
	fmt.Println(v, p)               //v is null and p is false
	for key, value := range m1 {
		_, _ = key, value //dosomething
	}
}

//M3.3.1 Structs - organises related data in groups
func structsTest() {
	type Data struct {
		name    int
		address int
		phone   int
	}

	var d1 Data
	d1.name = 1
	fmt.Println(d1)
	d2 := Data{1, 2, 3} //literal declaration, can also explicitly name with {field:value}
	_ = d2
}
