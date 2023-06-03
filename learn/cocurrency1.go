package learn

import (
	"fmt"
	"sync"
	"time"
)

/*
A goroutine is a "lightweight thread" managed by the Go runtime.

go f(x, y, z) // starts a new goroutine running

The evaluation of f, x, y, and z happens in the current goroutine
and the execution of f happens in the new goroutine.

Goroutines run in the same address space,
so access to shared memory must be synchronized!!

The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func RUN_0() {
	go say("world")
	say("hello")
}

/*
	導管

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and

	// assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

ch := make(chan int)
By default, sends and receives "block" until the other side is ready.
This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice,
distributing the work between two goroutines.

Once both goroutines have completed their computation, it calculates the final result.
*/
func sum(slice []int, channel chan int) {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	channel <- sum // send sum to channel
}

func RUN_1() {
	slice0 := []int{7, 2, 8, -9, 4, 0}

	channel0 := make(chan int)
	go sum(slice0[:len(slice0)/2], channel0)
	go sum(slice0[len(slice0)/2:], channel0)
	x, y := <-channel0, <-channel0 // receive from channel

	fmt.Println(x, y, x+y)
}

/*
Channels can be buffered.

Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.
*/

func RUN_2() {
	ch := make(chan int, 2)
	ch <- 3
	ch <- 2
	// ch <- 2 <-- this line cause deadlock, buffer is 2 int
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.


Note:
Only the sender should close a channel, never the receiver.
Sending on a closed channel will cause a panic.
Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/

func fibonacci(number int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < number; i++ {
		channel <- x
		x, y = y, x+y
	}
	close(channel)
}

func RUN_3() {
	bufferChan := make(chan int, 10)
	go fibonacci(cap(bufferChan), bufferChan)

	for i := range bufferChan {
		fmt.Println(i)
	}
}

/*
The select statement lets a goroutine "wait" on multiple communication operations.
A select "blocks until one of its cases can run", then it executes that case.
It chooses one at random if multiple are ready.
*/

func fibonacci2(channel, quitChan chan int) {
	x, y := 0, 1
	for {
		select {
		case channel <- x:
			x, y = y, x+y
		case <-quitChan:
			fmt.Println("quit")
			return
		}
	}
}

func RUN_4() {
	channel0 := make(chan int)
	quitChan := make(chan int)

	// a goroutine to receive from 2 channels
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel0)
		}
		quitChan <- 0
	}()

	// run to send stuff to 2 channels
	fibonacci2(channel0, quitChan)
}

/*
The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:

select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
*/

func RUN_5() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

/*
sync.Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication?
What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called mutual exclusion,
and the conventional name for the data structure that provides it is mutex.

Go's standard library provides mutual exclusion with sync.Mutex and its two methods:

Lock
Unlock

We can define a block of code to be executed in mutual exclusion
by surrounding it with a call to Lock and Unlock as shown on the Inc method.

We can also use defer to ensure the mutex will be unlocked as in the Value method.
*/

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mutex sync.Mutex
	value map[string]int
}

// Inc increments the counter for the given key.
func (counter *SafeCounter) Inc(key string) {
	counter.mutex.Lock()

	// Lock so only one goroutine at a time can access the map c.v.
	counter.value[key]++

	counter.mutex.Unlock()
}

// Value returns the current value of the counter for the given key.
func (counter *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map c.v.
	counter.mutex.Lock()

	// use defer to ensure the mutex will be unlocked
	defer counter.mutex.Unlock()

	return counter.value[key]
}

func RUN_6() {
	counterMap := SafeCounter{value: make(map[string]int)}

	for i := 0; i < 1000; i++ {
		go counterMap.Inc("someKey")
	}

	time.Sleep(time.Second)

	fmt.Println(counterMap.Value("someKey"))
}
