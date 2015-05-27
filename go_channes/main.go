package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int, 10)
	done := make(chan bool)

	go sendData(data, done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 20)
	merge(data, done)

}

func sendData(data chan int, done chan bool, numbers ...int) {
	defer close(data)
	defer close(done)
	for _, number := range numbers {
		if number == 20 {
			time.Sleep(2000 * time.Millisecond)
		}
		data <- number
	}
	done <- true

}

func merge(data chan int, done chan bool) {
	for {
		select {
		case number := <-data:
			fmt.Println(number)
		case <-done:
			fmt.Println("The End")
			return
		}
	}
}
