package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "send result")
			case <-done:
				fmt.Println(idx, "exiting")
			}
		}(i)
	}

	fmt.Println("result:", <-ch)
	close(done)

	time.Sleep(3 * time.Second)
}
