package main

import (
	"fmt"
)

func main() {
	var a = "initial"
	fmt.Println("a =", a)
	a = "new value"
	fmt.Println("a = ", a)

	b := "b value"
	fmt.Printf("b = %T\n", b)
	b = "b new val" // b = 23 it will show error because we declare b as a String
	fmt.Println("b = ", b)

	// assigning multiple values at a time
	var age, price int = 23, 10000
	fmt.Println("age =", age, "price = ", price)

	d, f, g := 2, 3.5, "4"
	fmt.Println(d, f, g)

	var r = true
	fmt.Println(r)

	fmt.Println("----------------------")
	// zero_value
	var x int
	fmt.Println(x)
	var y float32
	fmt.Println(y)
	var z bool
	fmt.Println(z)
	var s string
	fmt.Println(s)
}
