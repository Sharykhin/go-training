package main

import (
	"fmt"
)

func main() {
	var x int = 2

	fmt.Println(x)
	{

		fmt.Println(x)
		var x string = "hello"
		fmt.Println(x)
	}
	var a []int

	a = append(a, 1)
	fmt.Println(a)

	m := make(map[string]int)

	m["one"] = 1
	fmt.Println(m)

}
