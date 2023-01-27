// for is Go’s only looping construct. Here are some basic types of for loops.
package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println("i = ", i)
		i = i + 1
	}

	for j := 0; j < 11; j++ {
		fmt.Println("j = ", j)
		if j == 5 {
			break
		}

	}
	for {
		fmt.Println("loop break")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n = ", n)
	}

}
