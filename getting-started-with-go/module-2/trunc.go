package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Enter a float: ")
	fmt.Scan(&x)
	y := int(x)
	fmt.Println(y)
}
