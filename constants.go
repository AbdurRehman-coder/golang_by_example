// Constants in golang

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const a = 10000000
	fmt.Println(a)

	const b = 200
	const result = a + b
	fmt.Println(result)

	const x = 3e20 / a
	fmt.Println(x)

	fmt.Println(int64(x))

	fmt.Println(math.Sin(a))

}
