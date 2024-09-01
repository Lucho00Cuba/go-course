package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(sli []int, i int) {
	t := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = t
}

func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < len(sli); i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	var inp string
	var nums []int

	for i := 0; i < 10; i++ {

		fmt.Print("Enter an integer: ")
		fmt.Scan(&inp)

		if strings.ToLower(inp) == "x" {
			break
		}

		num, _ := strconv.Atoi(inp)
		nums = append(nums, num)
	}

	BubbleSort(nums)
	fmt.Println(nums)
}
