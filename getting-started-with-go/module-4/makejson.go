package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Address string
}

func main() {
	var name, address string

	fmt.Print("Enter a name: ")
	fmt.Scan(&name)

	fmt.Print("Enter an address: ")
	fmt.Scan(&address)

	p := Person{name, address}
	json, _ := json.Marshal(p)

	fmt.Println(string(json))
}
