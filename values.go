// values in Go

package main

import "fmt"

func main() {
	// String values
	fmt.Println("go " + "lang")

	// integer values
	fmt.Println("1 + 1 = ", 1+1)

	// float values
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// boolean values
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}

// Outputs:
go lang
1 + 1 =  2
7.0/3.0 =  2.3333333333333335
false
true
true
