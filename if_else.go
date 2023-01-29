// if else condition in go
// Note: There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} else {
		fmt.Println("8 is not divisible by 4")
	}

	if a := 9; a < 0 {
		fmt.Println(a, "is negative")
	} else if a < 10 {
		fmt.Println(a, "has 1 digit")
	} else {
		fmt.Println(a, "has multiple digits")
	}

}
