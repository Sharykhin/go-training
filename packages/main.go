package main

import (
	"fmt"
	_ "go-training/packages/init"
	"private/validation"
)

func main() {

	var validator validation.Validator

	if ok := validator.Requred(" "); ok {
		fmt.Println("Yes it is true")
	} else {
		fmt.Println("No it was false")
	}

}
