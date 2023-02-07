package main

import "fmt"

func main() {
	// Here we create an array a that will hold exactly 5 ints.
	// The type of elements and length are both part of the array’s type.
	//  By default an array is zero-valued, which for ints means 0s.
	var a [5]int
	fmt.Println("a: ", a)

	// We can set a value at an index using the array[index] = value syntax,
	// and get a value with array[index].
	a[4] = 1000
	fmt.Println("set a: ", a)
	fmt.Println("get a: ", a[4])

	fmt.Println("a len: ", len(a))

	b := [5]int{3, 4, 5, 6, 9}
	fmt.Println("b arr", b, "\nlength of b", len(b))

	k := [3]string{"Rehman", "Ishfaq", "Sheraz"}
	fmt.Println("k = ", k[1])
	// Array types are one-dimensional,
	// but you can compose types to build multi-dimensional data structures.
	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d array: ", twoD)
}
