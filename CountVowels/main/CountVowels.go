// Write a Go program to count the number of vowels in a string.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func contains(s []string, c string) bool {
	for _, d := range s {
		if d == c {
			return true
		}
	}
	return false
}

func main() {
	var num int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")
	text, _ := reader.ReadString('\n')
	var vowels = []string{"a", "e", "i", "o", "u"}
	for _, c := range text {
		if contains(vowels, string(c)) { // string(variable) is equivalent to str(variable) of python
			num++
		}
	}
	fmt.Println(num)
}
