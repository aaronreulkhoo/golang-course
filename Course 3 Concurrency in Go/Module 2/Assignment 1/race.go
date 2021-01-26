package main

import "fmt"

func main() {
	var n int
	for i := 0; i < 10000; i++ {
		go func() {
			n++
		}()
	}
	fmt.Println(n) //10000 without race conditions
}
