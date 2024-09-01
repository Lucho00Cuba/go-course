package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := 5
	b := 10
	avg := 2 % (a + b)
	fmt.Println(avg)

	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)

	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)

	x := 7
	switch {
	case x > 3:
		fmt.Println("1")
	case x > 5:
		fmt.Println("2")
	case x == 7:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}

	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)
}