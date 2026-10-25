package main

import "sync"

func sumPart(numbers []int, resultChannel chan<- int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	resultChannel <- sum
}

func ParallelSum(numbers []int) int {
	middle := len(numbers) / 2
	resultChannel := make(chan int, 2)

	go sumPart(numbers[:middle], resultChannel)
	go sumPart(numbers[middle:], resultChannel)

	firstSum := <-resultChannel
	secondSum := <-resultChannel

	return firstSum + secondSum
}

type SafeCounter struct {
	value int
	mutex sync.Mutex
}

func (counter *SafeCounter) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	counter.value++
}

func (counter *SafeCounter) Value() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()

	return counter.value
}

func RunCounter() int {
	const (
		goroutineCount      = 5
		incrementsPerWorker = 1000
	)

	counter := &SafeCounter{}
	var waitGroup sync.WaitGroup

	waitGroup.Add(goroutineCount)
	for worker := 0; worker < goroutineCount; worker++ {
		go func() {
			defer waitGroup.Done()

			for increment := 0; increment < incrementsPerWorker; increment++ {
				counter.Increment()
			}
		}()
	}

	waitGroup.Wait()
	return counter.Value()
}
