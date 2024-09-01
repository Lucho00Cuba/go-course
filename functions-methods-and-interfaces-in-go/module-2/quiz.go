package main

import "fmt"

var f func(string) int

func test(x string) int {
	return len(x)
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	f = test
	fmt.Println(f("hello"))

	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
}
