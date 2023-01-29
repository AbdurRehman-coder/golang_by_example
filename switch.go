// Switch statement in golang

package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("write ", i, " as ")
	// switch statement on i
	// Note: we don't use paranthesis in switch statement in go
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same case statement.
	// We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	// switch without an expression is an alternate way to express if/else logic.
	// Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	// A type switch compares types instead of values. You can use this to discover the type of an interface value.
	//  In this example, the variable t will have the type corresponding to its clause.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatAmI(14397293479734)
	whatAmI(false)
	whatAmI(343423423423104878379472974987394798274987398749374983479837498739424324234243234234234.3)
	whatAmI("out")

}
