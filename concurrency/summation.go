package concurrency

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

func Sum(subArrA []int) int {
	if len(subArrA) == 0 {
		return 0
	}
	if len(subArrA) == 1 {
		return subArrA[0]
	}
	log.Debug(subArrA[0], subArrA[1], subArrA)
	return subArrA[0] + subArrA[1] + Sum(subArrA[2:])
}

func sumWorker(wg *sync.WaitGroup, subArr []int, result chan int) {
	defer wg.Done()
	result <- Sum(subArr)
}

func ConcurrentSummation() {
	var wg sync.WaitGroup
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numOfWorkers := 1
	result := make(chan int, numOfWorkers)
	multiplier := len(arr) / numOfWorkers
	for i := 0; i < numOfWorkers; i++ {

		wg.Add(1)
		go func(start, end int) {
			log.Debug(start, end, len(arr), multiplier)
			if end > len(arr) || (len(arr)-end) < multiplier {
				end = len(arr)
			}
			if start > end || start > len(arr) {
				start = end
			}
			log.Debug(start, end, len(arr), multiplier)
			sumWorker(&wg, arr[start:end], result)
		}(i*multiplier, (i+1)*multiplier)
	}
	wg.Wait()
	sumOfInts := 0
	for i := 0; i < numOfWorkers; i++ {
		sumOfInts += <-result
	}
	close(result)
	log.Debugf("\nFinal Sum of the data : %v ", sumOfInts)
}
