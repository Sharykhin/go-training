package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, done, &wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(wockerId int, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", wockerId)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", wockerId)
}
