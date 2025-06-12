package main

import "fmt"

func main() {
	// Declare an empty value array
	var empty_arr [5]int
	fmt.Println("Array with empty values:", empty_arr)
	fmt.Println(len(empty_arr))
	fmt.Println(cap(empty_arr)) // capacity

	// Declare an empty array
	var arr0 = []int{}
	fmt.Println("Empty Array :", arr0)
	fmt.Println(len(arr0))
	fmt.Println(cap(arr0)) // capacity

	// Declare an array with initial values
	var arr1 = [5]int{1, 2, 3, 4}
	fmt.Println("Array with initial values:", arr1)
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1)) // capacity

	// Declare an array with a shorthand notation
	var arr2 = [...]int{6, 7, 8, 9, 10} // ... means the compiler will determine the size of the array
	fmt.Println("Array declared using shorthand notation:", arr2)
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2)) // capacity

}
