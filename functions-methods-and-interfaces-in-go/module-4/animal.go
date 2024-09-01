package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	sound      string
}

type Animals interface {
	Eat()
	Move()
	Speak()
	Command(action string)
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func (a Animal) Command(action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func main() {
	m := make(map[string]Animal)
	m["cow"] = Animal{"grass", "walk", "moo"}
	m["bird"] = Animal{"worms", "fly", "peep"}
	m["snake"] = Animal{"mice", "slither", "hsss"}

	var a Animals

	for {
		var command, animal, action string

		fmt.Print("> ")
		fmt.Scan(&command, &animal, &action)

		if command == "query" {
			a = m[animal]
			a.Command(action)
		} else if command == "newanimal" {
			m[animal] = m[action]
			fmt.Println("Created it!")
		}
	}
}
