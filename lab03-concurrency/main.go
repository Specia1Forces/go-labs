package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	sum := ParallelSum(numbers)
	fmt.Println("Parallel sum:", sum)

	counterValue := RunCounter()
	fmt.Println("Safe counter:", counterValue)
}
