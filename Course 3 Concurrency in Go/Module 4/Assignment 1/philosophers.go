package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	left, right *Chopstick
}

func (p Philosopher) eat(hostch chan bool, wg *sync.WaitGroup, i int) {
	for e := 0; e < 3; e++ {
		<-hostch
		p.left.Lock()
		p.right.Lock()
		fmt.Println("Starting to eat", i)
		fmt.Println("Eating...")
		fmt.Println("Finishing eating", i)
		p.left.Unlock()
		p.right.Unlock()
		hostch <- true
	}
	wg.Done()
}

func host(hostch chan bool, wg *sync.WaitGroup, maxEating int) {
	for i := 0; i < maxEating; i++ {
		hostch <- true
	}
}

func main() {
	maxEating := 2
	hostch := make(chan bool, maxEating)
	var wg sync.WaitGroup
	numPhilosophers := 5
	wg.Add(numPhilosophers)
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}
	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{chopsticks[i], chopsticks[(i+1)%numPhilosophers]}
	}
	go host(hostch, &wg, maxEating)
	for i := 0; i < numPhilosophers; i++ {
		go philosophers[i].eat(hostch, &wg, i+1)
	}
	wg.Wait()
}
