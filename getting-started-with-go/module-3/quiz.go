package main

import "fmt"

type P struct {
	x string
	y int
}

func main() {
	x := []int{4, 8, 5}
	y := -1
	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}
	fmt.Println(y)

	x2 := [...]int{4, 8, 5}
	y2 := x2[0:2]
	z := x2[1:3]
	y2[0] = 1
	z[1] = 3
	fmt.Println(x2)

	x3 := [...]int{1, 2, 3, 4, 5}
	y3 := x3[0:2]
	z3 := x3[1:4]
	fmt.Println(len(y3), cap(y3), len(z3), cap(z3))

	x4 := map[string]int{
		"ian": 1, "harris": 2}
	for i, j := range x4 {
		if i == "harris" {
			fmt.Println(i, j)
		}
	}

	b := P{"x", -1}
	a := [...]P{P{"a", 10},
		P{"b", 2},
		P{"c", 3}}
	for _, z5 := range a {
		if z5.y > b.y {
			b = z5
		}
	}
	fmt.Println(b.x)

	s := make([]int, 0, 3)
	s = append(s, 100)
	fmt.Println(len(s), cap(s))
}
