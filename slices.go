package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	// We can set and get just like with arrays.
	s := make([]string, 3)
	fmt.Println("emps: ", s)

	// add value to slices
	s[0] = "Rehman"
	s[1] = "Ujala"
	fmt.Println("emps: ", s[1], "length: ", len(s)) // length is 3, but values are two inside, len defined it start

	// append new values
	s = append(s, "Haris")
	s = append(s, "Salma", "Zuhra", "Afreen")
	fmt.Println("apd: ", s, "len: ", len(s))

	// Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s)) // create new slices of name 'c' and give same size of 's'
	// now copy the same size slices
	copy(c, s)
	fmt.Println("c copy: ", c)

	l := s[2:5]
	fmt.Println("l = ", l)
	l = s[2:]
	fmt.Println("l = ", l)
	l = s[:5]
	fmt.Println("l = ", l)

	t := []string{"i", "k", "j", "z"}
	fmt.Println("t ", reflect.TypeOf(t))

	// Slices can be composed into multi-dimensional data structures.
	//  The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
		fmt.Println("twoD:", twoD)
	}

}
