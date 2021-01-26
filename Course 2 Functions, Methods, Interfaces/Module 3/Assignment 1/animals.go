package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal : Struct to contain all animal data
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat : prints the associated Animal's food
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

// Move : prints the associated Animal's locomotion
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

// Speak : prints the associated Animal's noise
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	stdin := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter an animal and an action seperated by a space, or leave empty to exit.")
	flag := true
	for flag {
		fmt.Printf(">")
		scanned, _ := stdin.ReadString('\n')
		scanned = strings.Replace(scanned, "\r", "", -1)
		scanned = strings.Replace(scanned, "\n", "", -1)
		if scanned == "" {
			fmt.Println("Exited.")
			flag = false
		}
		sli := strings.Split(scanned, " ")
		if len(sli) != 2 {
			fmt.Println("Wrong number of arguments! Did you leave a trailing space?")
		} else if val, ok := animals[sli[0]]; !ok {
			fmt.Println("Animal not found!")
		} else {
			switch sli[1] {
			case "eat":
				val.Eat()
			case "move":
				val.Move()
			case "speak":
				val.Speak()
			default:
				fmt.Println("Animal action not supported!")
			}
		}
		// if val,ok:=animals[sli]

	}

}
