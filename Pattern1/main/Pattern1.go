// *
// * *
// * * *
// * * * *
// * * * * *

package main

import "fmt"

func repeat(str1 string, number int) {
	for i := 1; i <= number; i++ {
		fmt.Print(str1)
	}
	return
}

func main() {
	for j := 1; j < 6; j++ {
		repeat("* ", j)
		fmt.Println()
	}
}
