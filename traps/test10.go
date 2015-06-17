package main

import (
	"fmt"
)

func main() {

loop:
	for {
		switch {
		case true:
			fmt.Println("yes we've got it")
			break loop
		}
	}

	fmt.Println("Thanks, we found an exit")
}
