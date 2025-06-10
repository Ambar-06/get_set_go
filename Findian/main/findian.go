// Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

// Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

// Submit your source code for the program,
// “findian.go”.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var randomString string
	fmt.Println("Please enter a string:")
	reader := bufio.NewReader(os.Stdin)
	randomString, _ = reader.ReadString('\n')
	randomString = strings.ToLower(strings.TrimSpace(randomString))
	if strings.HasPrefix(randomString, "i") && strings.HasSuffix(randomString, "n") && strings.Contains(randomString, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
