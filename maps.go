package main

import (
	"fmt"
	"reflect"
)

func main() {

	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	var m = make(map[string]int)
	fmt.Println("m: ", reflect.TypeOf(m))

	m["k1"] = 7
	m["k2"] = 129
	fmt.Println("m: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("m: ", m)

	_, prs := m["K2"]
	fmt.Println("prs: ", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo1": 23, "foo2": 45}
	fmt.Println("n = ", n)

}
