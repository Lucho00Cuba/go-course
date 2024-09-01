package main

import (
	"fmt"
	"time"
)

var count int

func mult() {
	fmt.Println("2 x", count)
	count *= 2
}

func incr() {
	fmt.Println("1 +", count)
	count++
}

func main() {
	count = 1

	go mult()
	go mult()

	go incr()
	go incr()

	time.Sleep(500 * time.Millisecond)

	fmt.Println("  =", count)
}

/*
	Four goroutines change one global variable.
	Two of the routines will double the variable
	and the other two will increment the variable.
	The initial starting value of the global variable is 1
	and each go routine will print the count as it sees.
	The maximum value will be 12; however, 9 is much more likely.
	The final counts are dependent on the goroutine order as
	1*2+2 differs from 1+2*2.
*/
