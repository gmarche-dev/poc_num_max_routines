package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const totalLength = 2000000
	var wg sync.WaitGroup
	wg.Add(totalLength)
	ch := make(chan int, totalLength)
	for i := 0; i < totalLength; i++ {
		go func() {
			<-time.After(5 *time.Second)
				ch <- 1
				wg.Done()
				return
		}()
	}
	wg.Wait()
	close(ch)
	fmt.Println("*************")
	fmt.Println(fmt.Sprintf("number of row expected : %d | number of row send %d", totalLength, len(ch)))
	fmt.Println("*************")
}