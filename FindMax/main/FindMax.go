// Write a Go program to find the largest of three numbers.

package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("Enter first number:")
	fmt.Scan(&a)
	fmt.Println("Enter second number:")
	fmt.Scan(&b)
	fmt.Println("Enter third number:")
	fmt.Scan(&c)

	if a == b && b == c {
		fmt.Println("All are equal")
	} else if a > b && a > c {
		fmt.Println("a is the largest")
	} else if b > a && b > c {
		fmt.Println("b is the largest")
	} else {
		fmt.Println("c is the largest")
	}
}
