package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	FirtsName []byte
	LastName  []byte
}

func main() {
	var path string
	names := make([]Name, 0)

	fmt.Print("Enter file path: ")
	fmt.Scan(&path)

	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		fname := []byte(line[0])
		lname := []byte(line[1])

		if len(fname) > 20 {
			fname = fname[:20]
		}
		if len(lname) > 20 {
			lname = lname[:20]
		}

		name := Name{fname, lname}
		names = append(names, name)
	}

	for _, n := range names {
		fmt.Printf("Firts Name: %s Last Name: %s\n", string(n.FirtsName), string(n.LastName))
	}
}
