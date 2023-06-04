package learn

import (
	"fmt"
	"sync"
)

type TaskResult struct {
	ID     int
	Result int
}

/*
Here, we have a common pattern:
1. the worker have 2 arguments:
- input receive channel
- output receive channel
*/
func worker4(
	id int,
	tasks <-chan int,
	results chan<- TaskResult,
	wg *sync.WaitGroup,
	mutex *sync.Mutex) {

	defer wg.Done()

	for num := range tasks {
		// Perform the task
		result := num * num

		// Store the result
		mutex.Lock()
		results <- TaskResult{ID: id, Result: result}
		mutex.Unlock()
	}
}

func RUN_CO4() {
	numTasks := 10
	numOfWorkers := 4

	// Create channels for tasks and results
	tasks := make(chan int)
	results := make(chan TaskResult)

	// Create a mutex to protect the results
	mutex := &sync.Mutex{}

	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup

	// Start the worker goroutines
	for id := 1; id <= numOfWorkers; id++ {
		wg.Add(1)
		go worker4(id, tasks, results, &wg, mutex)
	}

	// Start the task goroutines via self-invoke fn
	go func() {
		for i := 1; i <= numTasks; i++ {
			tasks <- i
		}
		close(tasks) // close the tasks channel after all tasks are done
	}()

	// Collect the results from the worker goroutines
	go func() {
		wg.Wait()
		close(results) // close the results channel after results channel are done
	}()

	// Process the results
	for result := range results {
		fmt.Printf("Worker %d: Task %d Result %d\n", result.ID, result.ID, result.Result)
	}
}
