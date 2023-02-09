package main

import (
	"fmt"
	"reflect"
)

// range iterates over elements in a variety of data structures.
//  Let’s see how to use range with some of the data structures we’ve already learned.

func main() {
	nums := []int{1, 3, 7}
	fmt.Println("nums: ", nums, reflect.TypeOf(nums))

	sum := 0
	for _, num := range nums {
		sum += num
		fmt.Println(num)
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		fmt.Println("i = ", i)
		if num == 3 {
			fmt.Println("index = ", i)
		}
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// range through map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("keys: ", k)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
