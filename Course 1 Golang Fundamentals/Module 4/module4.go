package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// jsonTest()
	ioutilTest()
	osFileTest()
}

// Protocols and Formats
// M4.1.1 RFCs - Request for Comments - standards for communications
// http.Get("www.uci.edu")
// net.Dial("tcp", "uci.edu:80")

// e.g. of protocols are HTML, URI, HTTP, etc
//Golang provides packages for important RFCs to encode and decode protocol format

// M4.1.2 JSON - RFC7159, format to represent structured information > similar to map/struct
//JSON marshalling is the act of generating a json representation from an object
/* NOTE: Fields have to start with a capital letter or marshalling will return empty byte array */
type Data struct {
	Field1 int
	Field2 int
	Field3 int
}

func jsonTest() {
	d1 := Data{1, 2, 3}
	var d2 Data
	byteArray, err := json.Marshal(d1) //returns json representation (byte array) of struct
	fmt.Println(string(byteArray))
	fmt.Println(d1, d2, err)                     //different
	err = json.Unmarshal([]byte(byteArray), &d2) //returns struct from json object, needs to fit json fields
	fmt.Println(d1, d2, err)                     //same
}

// M4.2.1 File Access via ioutil - basic open, read, write, close, seek
func ioutilTest() {
	var hello string = "this was written by ioutiltest"
	err := ioutil.WriteFile("test.txt", []byte(hello), 0777) //creates files, writes []byte, closes - uses unix style permissions
	data, err := ioutil.ReadFile("test.txt")                 //open, read, and close, however large files cause substantial delays
	fmt.Println(string(data), err)
}

// M4.2.2 File Access via os - provides more specific control
func osFileTest() {
	f, _ := os.Create("os.txt")     // opens a file, returns file descriptor
	nb, _ := f.WriteString("hello") //writes a string
	bArr := []byte("abc")
	nb, _ = f.Write(bArr) //writes a []byte
	f.Close()
	f, _ = os.Open("os.txt") // opens a file, returns file descriptor
	bArr = make([]byte, 10)
	nb, err := f.Read(bArr) // reads a file filling the provided []byte, returns number of bytes read
	fmt.Println(string(bArr), nb, err)
	f.Close() // closes a file

}
