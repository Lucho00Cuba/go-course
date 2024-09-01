package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Print("Enter the\n")
	fmt.Print("  acceleration: ")
	fmt.Scan(&a)

	fmt.Print("  initial velocity: ")
	fmt.Scan(&v0)

	fmt.Print("  initial displacement: ")
	fmt.Scan(&s0)

	s := GenDisplaceFn(a, v0, s0)

	fmt.Print("  time: ")
	fmt.Scan(&t)

	fmt.Println(s(t))
}
