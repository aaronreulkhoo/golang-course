package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// rangeChannelTest()
	// selectTest()
	// mutexTest()
	// onceTest()
	philosopherTableTest()
}

//M4.1.1 Blocking on Channels
/* It is common to iterate on channels -> use the range function to continously read from channel
Use close(channel) to close the channel and end the loop*/
func rangeChannelTest() {
	c := make(chan int)
	go func(c chan int) {
		c <- 6
	}(c)
	for i := range c {
		fmt.Println(i)
		close(c)
	}
}

//M4.1.2 Select Statement to choose from multiple channels
/* Unlike waiting on two channels which may end up with 1 blocking, select executes whichever is unblocked first
May select either send or receive operations
Typically used with an abort channel, send anything to return and exit the loop
May also want to include a default operation to avoid blocking*/
func selectTest() {
	inchan, outchan, abort := make(chan int), make(chan int), make(chan int)
	go func(c chan int) {
		c <- 6
	}(inchan)
	go func(c chan int) {
		<-c
	}(outchan)
	go func(c chan int) {
		time.Sleep(1000)
		c <- 404
	}(abort)
	for {
		select {
		case a := <-inchan:
			fmt.Printf("received %d from channel a\n", a)
		case outchan <- 6:
			fmt.Println("sent 6 to channel b")
		case <-abort:
			fmt.Println("aborted")
			return
		default:
			fmt.Println("Nothing")
		}
	}
}

//M4.2.1 Mutual Exclusion of Threads
/* Sharing variables concurrently may cause problems unless they are concurrency safe (invoked conrurrently without interfering with others)
Concurrency problems occur at the machine code level (which Go is compiled to), interleavings
*/

//M4.2.2 Mutexes in Go
/* sync.Mutex employs a binary semaphore */
func mutexTest() {
	var mtx sync.Mutex
	x := 0
	go func() {
		mtx.Lock()
		x++
		fmt.Println("Releasing mutex from 1")
		mtx.Unlock()
	}()
	go func() {
		mtx.Lock()
		x++
		fmt.Println("Releasing mutex from 2")
		mtx.Unlock()
	}()
	time.Sleep(1000)
	fmt.Println(x)
}

//M4.3.1 Synchronous Initialization
/* sync.Once provides a one-time mutex which has a method Do(f) which performs the function f one time
All other calls to once.Do() are blocked - can be used to ensure initialization in a subroutine*/
func onceTest() {
	var wg sync.WaitGroup
	go goRoutine(&wg)
	wg.Add(1)
	wg.Wait()
	fmt.Println("last")
}

func goRoutine(wg *sync.WaitGroup) {
	var once sync.Once
	once.Do(func() {
		fmt.Println("init")
	})
	fmt.Println("hello")
	wg.Done()
}

//M3.4.2 Deadlock
/* Deadlock comes from circular synchronisation dependencies
e.g. goroutine1 ch<-1 and mtx.Unlock(), goroutine 2 x:=<-ch and mtx.Lock()
Go automatically detects when all routines are deadlocked
Below is example of the classic philosophers table*/
type Utensil struct {
	sync.Mutex
}

type Philosopher struct {
	left, right *Utensil
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	p.left.Lock()
	p.right.Lock()
	fmt.Println("Eating")
	p.left.Unlock()
	p.right.Unlock()
	wg.Done()
}
func philosopherTableTest() {
	for x := 0; x < 100; x++ {
		var wg sync.WaitGroup
		utensils := make([]*Utensil, 5)
		for i := 0; i < 5; i++ {
			utensils[i] = new(Utensil)
		}
		philosophers := make([]*Philosopher, 5)
		for i := 0; i < 5; i++ {
			philosophers[i] = &Philosopher{utensils[i], utensils[(i+1)%5]}
		}
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go philosophers[i].eat(&wg)
		}
		wg.Wait()
	}

}
