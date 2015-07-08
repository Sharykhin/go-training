package main

import (
	"fmt"
)

type Human struct {
	name string
}

func main() {
	var num int = 10
	Peter := Human{name: "Peter"}

	fmt.Println(num)
	fmt.Println(Peter)
}
