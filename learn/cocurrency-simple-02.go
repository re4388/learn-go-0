package learn

import (
	"fmt"
)

func sum2(numbers []int, result chan<- int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	// this is a special thing in Go
	// we stor the result back into the passing in argument
	// this is common is this argument is a receiver side of channel
	result <- sum
}

func RUN_CO2() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// We create a channel named result using the make function to receive the sum result.
	result := make(chan int)

	go sum2(numbers[:len(numbers)/2], result) // Calculate sum of first half of numbers concurrently
	go sum2(numbers[len(numbers)/2:], result) // Calculate sum of second half of numbers concurrently

	sum1 := <-result // Receive the first sum from the channel
	sum2 := <-result // Receive the second sum from the channel

	totalSum := sum1 + sum2
	fmt.Println("Total sum:", totalSum)
}
