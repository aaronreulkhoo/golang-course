package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Animal interface
type Animal interface {
	eat()
	move()
	speak()
}

// Cow type
type Cow struct{ food, locomotion, noise string }

func (cow Cow) eat()   { fmt.Println(cow.food) }
func (cow Cow) move()  { fmt.Println(cow.locomotion) }
func (cow Cow) speak() { fmt.Println(cow.noise) }

// Bird type
type Bird struct{ food, locomotion, noise string }

func (bird Bird) eat()   { fmt.Println(bird.food) }
func (bird Bird) move()  { fmt.Println(bird.locomotion) }
func (bird Bird) speak() { fmt.Println(bird.noise) }

// Snake type
type Snake struct{ food, locomotion, noise string }

func (snake Snake) eat()   { fmt.Println(snake.food) }
func (snake Snake) move()  { fmt.Println(snake.locomotion) }
func (snake Snake) speak() { fmt.Println(snake.noise) }

func main() {
	animals := map[string]Animal{}
	stdin := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a command or type \"help\" for a list of commands.")
	for {
		fmt.Printf("createAnimals.go>")
		scanned, _ := stdin.ReadString('\n')
		scanned = strings.Replace(scanned, "\r", "", -1)
		scanned = strings.Replace(scanned, "\n", "", -1)
		if scanned == "exit" {
			break
		} else if scanned == "help" {
			fmt.Printf("%-20s%s\n", "help", "Displays all available commands")
			fmt.Printf("%-20s%s\n", "ls", "Displays a list of all created animals.")
			fmt.Printf("%-20s%s\n", "newanimal", "Creates a new animal of the name and type - usage \"newanimal [name] [type]\".")
			fmt.Printf("%-20s%s\n", "query", "Queries an named animal for an action - usage \"query [name] [action]\".")
			fmt.Printf("%-20s%s\n", "exit", "Exits the program.")
			continue
		} else if scanned == "ls" {
			keys := reflect.ValueOf(animals).MapKeys()
			fmt.Println(keys)
			continue
		}
		sli := strings.Split(scanned, " ")
		if len(sli) != 3 {
			fmt.Println("Wrong number of arguments! Did you leave a trailing space?")
			continue
		} else if sli[0] == "newanimal" {
			if _, ok := animals[sli[1]]; ok {
				fmt.Printf("%s already exists!\n", sli[1])
				continue
			}
			switch sli[2] {
			case "cow":
				animals[sli[1]] = Cow{"grass", "walk", "moo"}
			case "bird":
				animals[sli[1]] = Bird{"worms", "fly", "peep"}
			case "snake":
				animals[sli[1]] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Println("Animal not supported!")
				continue
			}
			fmt.Printf("Created a %s called \"%s\"!\n", sli[2], sli[1])
		} else if sli[0] == "query" {
			if animal, ok := animals[sli[1]]; !ok {
				fmt.Println("Animal not yet created!")
				continue
			} else {
				switch sli[2] {
				case "eat":
					animal.eat()
				case "move":
					animal.move()
				case "speak":
					animal.speak()
				default:
					fmt.Println("Animal action not supported!")
				}
			}
		} else {
			fmt.Println("Command not recognised!")
		}

	}

}
