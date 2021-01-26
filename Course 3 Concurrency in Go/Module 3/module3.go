package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// goTest()
	// sleepTest()
	// syncTest()
	channelTest()
	syncTest2()
	bufferedChannelTest()
}

//M3.1.1 - Goroutines
/*Go has simple implementation of concurrency using the keyword go
foo() -> this is blocking
go foo() -> this is non blocking
A goroutine exits when its code is complete, however the main goroutine forces termination of all current goroutines
Thus, a goroutine may not actually complete execution if the main completes early
*/
func goTest() {
	go fmt.Println("go test") //not printed
	fmt.Println("test")
}

//M3.1.2 - Exiting Goroutines
/* As early main exits cause goroutines to terminate, can delay using time.Sleep(t), however it relies on assumption (bad practice)
 */
func sleepTest() {
	go fmt.Println("go test") //no printed
	time.Sleep(1000)
	fmt.Println("test")
}

//M3.1.2 - Basic Synchronization
/* As early main exits cause goroutines to terminate, can delay using time.Sleep(t), however it relies on assumption (bad practice)
Goroutines are large independent have no understanding of timing going on in other goroutines
Synchronisations use global events which is viewed simultaneously by all goroutine threads, restricts bad interleavings
*/

//M3.2.1 - Wait Groups
/* sync.WaitGroup forces goroutines to wait for all others in the same group
Contains an internal counter that increments for each goroutine, decrements for each complete goroutine, and blocks when not 0
*/
func syncTest() {
	var wg sync.WaitGroup
	wg.Add(1) //increments
	go goRoutine(&wg)
	wg.Wait() //blocks
}

func goRoutine(wg *sync.WaitGroup) {
	fmt.Println("goRoutine")
	defer wg.Done() //decrements
}

//M3.3.1 - Communication Channels
/* Goroutines usually work together to perform a larger tasks
Channels transfer typed data, send and receive using arrow operator
*/
func channelTest() {
	ch := make(chan int)
	go multiply(1, 2, ch)
	go multiply(3, 4, ch)
	a := <-ch
	b := <-ch
	fmt.Println(a * b)
}

func multiply(v1 int, v2 int, c chan int) {
	c <- v1 * v2
}

//M3.3.2 - Blocking on Channels
/* Default channels are unbuffered and cannot hold data in transit
Sending to a channel blocks until data is received, receiving blocks until data is sent
E.g. Task 1 c<-3, Task 2 x:=<-3, so task 2 is blocked until task 1 sends
Thus, is also synchronizing and can be used by itself(throw away result) to achieve same effect as waitgroup
*/
func syncTest2() {
	var ch chan int
	go goRoutine2(ch)
	<-ch
}

func goRoutine2(ch chan int) {
	fmt.Println("goRoutine")
	ch <- 1
}

//M3.3.3 - Buffered Channels
/* Buffered channels can contain a limited nubmer of objects
Capacity is the number of objects it can hold in transit, passed as argument to make functions
Sending only blocks if buffer is full, receiving only blocks if buffer is empty
Buffering allows decoupling of sender and receiver (producer consumer paradigm)
*/

func bufferedChannelTest() {
	ch := make(chan int, 3)
	go multiply(1, 2, ch)
	go multiply(3, 4, ch)
	go multiply(5, 6, ch)
	a := <-ch
	b := <-ch
	c := <-ch
	fmt.Println(a * b * c)
}
