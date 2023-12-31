package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func SequentialPrint() {
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\nFinished sequential operation in : %v\n", time.Since(startTime))
}

func ConcurrentPrint() {
	var wg sync.WaitGroup
	resultChan := make(chan int, 12)
	// defer close(resultChan)
	sTime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(temp int) {
			defer wg.Done()
			resultChan <- temp
			time.Sleep(1 * time.Second)
		}(i)
	}
	go func() {
		wg.Wait()
		fmt.Println("Wait over")
		close(resultChan)
	}()
	for ch := range resultChan {
		fmt.Printf("Concurrent Value : %v\n", ch)
	}

	// <-resultChan

	fmt.Printf("\nFinished Concurrent operation in : %v\n", time.Since(sTime))
}
