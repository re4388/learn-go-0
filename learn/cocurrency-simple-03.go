package learn

import (
	"fmt"
	"sync"
)

// a counter with lock to ensure no other thread can access
// at the same time
type Counter struct {
	value int
	mutex sync.Mutex
}

// impl Increment for Counter type
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// *sync.WaitGroup.Done() is to wait all threads to finish
// worker fn take a counter ptr and wg ptr
func worker(counter *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter.Increment()
	}
}

func RUN_CO3() {
	counter := Counter{value: 0}

	// we use waitGroup to ensure all worker are done
	var wg sync.WaitGroup

	// Create 5 worker goroutines
	for i := 0; i < 5; i++ {
		// for each worker we setup, we add the worker into the waiting group
		wg.Add(1)
		go worker(&counter, &wg)
	}

	// Wait for all worker goroutines to finish
	wg.Wait()

	fmt.Println("Final value:", counter.value)

}
