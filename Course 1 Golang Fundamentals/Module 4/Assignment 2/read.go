package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type name struct {
	fname string
	lname string
}

func main() {
	ns := []name{}
	var filename string
	for {
		fmt.Printf("Please enter a file name: ")
		fmt.Scan(&filename)
		f, err := os.Open(filename)
		if err == nil {
			f.Close()
			break
		} else {
			fmt.Printf("Error: ")
			fmt.Println(err)
		}
	}
	f, _ := os.Open(filename)
	br := bufio.NewReader(f)
	var fname, lname string
	for {
		ln, _, _ := br.ReadLine()
		if ln == nil {
			break
		}
		s := strings.Split(string(ln), " ")
		if utf8.RuneCountInString(s[0]) >= 20 {
			fname = s[0][:20]
		} else {
			fname = s[0]
		}
		if utf8.RuneCountInString(s[1]) >= 20 {
			lname = s[1][:20]
		} else {
			lname = s[1]
		}

		ns = append(ns, name{fname, lname})
	}

	f.Close()
	for _, name := range ns {
		fmt.Printf("fname: %s, ", name.fname)
		fmt.Printf("lname: %s\n", name.lname)
	}
}
