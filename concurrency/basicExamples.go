package concurrency

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func Task(num int, workerNumber int) {
	log.Debugf("worker %v : Number  = %v\n", workerNumber+1, num+1)
}

func Worker(wg *sync.WaitGroup, tasks chan int, workerNumber int) {
	wg.Done()
	for task := range tasks {
		Task(task, workerNumber)
	}
}

func BasicUsagePrint() {
	var wg sync.WaitGroup
	tasks := make(chan int, 5)
	numberOfWorkers := 4
	startTime := time.Now()

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			Worker(&wg, tasks, i)
		}(i)
	}
	for i := 0; i < 5; i++ {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
	log.Debugf("\nTime took to complete the operation %v\n", time.Since(startTime))
}
