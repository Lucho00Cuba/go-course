package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
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

func sort(wait *sync.WaitGroup, sli []int) {
	fmt.Println(sli)
	BubbleSort(sli)
	wait.Done()
}

func inij(i *int, j *int, list1 []int, list2 []int, list3 []int, list4 []int) {
	if len(list1) != 0 {
		*i = list1[0]
		*j = 1
	}
	if len(list2) != 0 {
		if *j == 0 {
			*i = list2[0]
			*j = 2
		} else {
			if *i > list2[0] {
				*i = list2[0]
				*j = 2
			}
		}
	}
	if len(list3) != 0 {
		if *j == 0 {
			*i = list3[0]
			*j = 3
		} else {
			if *i > list3[0] {
				*i = list3[0]
				*j = 3
			}
		}
	}
	if len(list4) != 0 {
		if *j == 0 {
			*i = list4[0]
			*j = 4
		} else {
			if *i > list4[0] {
				*i = list4[0]
				*j = 4
			}
		}
	}
}

func read(nums *[]int) {
	reader := bufio.NewReader(os.Stdin)
	inp, _, _ := reader.ReadLine()
	inpNums := strings.Split(string(inp), " ")

	for _, s := range inpNums {
		n, _ := strconv.Atoi(s)
		*nums = append(*nums, n)
	}
}

func main() {
	var nums []int
	var wait sync.WaitGroup
	var sorts []int

	fmt.Println("Enter unsorted integers: ")

	read(&nums)

	size := len(nums) / 4
	for i := 0; i < 4; i++ {
		wait.Add(1)
		if i != 3 {
			go sort(&wait, nums[i*size:(i+1)*size])
		} else {
			go sort(&wait, nums[i*size:])
		}
	}
	wait.Wait()

	list1 := nums[:1*size]
	list2 := nums[1*size : 2*size]
	list3 := nums[2*size : 3*size]
	list4 := nums[3*size:]
	for t := 0; t < len(nums); t++ {
		var i, j int

		inij(&i, &j, list1, list2, list3, list4)
		sorts = append(sorts, i)

		switch j {
		case 1:
			list1 = list1[1:]
		case 2:
			list2 = list2[1:]
		case 3:
			list3 = list3[1:]
		case 4:
			list4 = list4[1:]
		}
	}
	fmt.Println(sorts)
}
