package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a string: ")
	scanner.Scan()

	t = scanner.Text()
	s := strings.ToLower(t)

	found := "Not Found!"
	if strings.HasPrefix(s, "i") {
		if strings.HasSuffix(s, "n") {
			if strings.Contains(s, "a") {
				found = "Found!"
			}
		}
	}
	fmt.Println(found)
}