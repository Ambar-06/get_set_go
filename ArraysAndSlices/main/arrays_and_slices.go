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

	// Slice : Slices are just like windows over the array, see the example below
	var new_ar = [...]string{"a", "b", "c", "d"} // declare an array of strings a,b,c,d
	s1 := new_ar[1:3]                            // [include:exclude] --  // s1 is a slice which only shows the element inside the window (In this case the window is represented as ()) ["a", ("b", "c"), "d"]
	fmt.Println("Slice1:", s1)
	fmt.Println(len(s1)) // 2

	fmt.Println(cap(s1)) // 3 Why 3?
	// cap() gives the capacity of the slice, which is:
	// The number of elements from the starting index of the slice (index 1) to the end of the original array.
	// In this case: from "b" (index 1) to "d" (index 3) — total 3 elements: "b", "c", "d" → so cap(s1) == 3

	// Slices can also be created directly without creating an array like
	sli1 := make([]int, 10) // make function arguments ; type, length/capacity ([]int, 10) or type, length, capacity ([]int, 10, 15)
	fmt.Println("Slice created directly: ", sli1)
	fmt.Println(len(sli1)) // 10
	fmt.Println(cap(sli1)) // 10

}
