package learn

/*
Go's approach to concurrency differs from the traditional use of threads and shared memory.
Philosophically, it can be summarized:

Don't communicate by sharing memory; share memory by communicating.

Channels allow you to pass references to data structures between goroutines.
If you consider this as passing around ownership of the data (the ability to read and write it),
they become a powerful and expressive synchronization mechanism.

In this codewalk we will look at a simple program that polls a list of URLs,
checking their HTTP response codes and periodically printing their state.
*/

import (
	"log"
	"net/http"
	"time"
)

const (
	numPollers     = 2                // number of Poller goroutines to launch
	pollInterval   = 60 * time.Second // how often to poll each URL
	statusInterval = 10 * time.Second // how often to log status to stdout
	errTimeout     = 10 * time.Second // back-off timeout on error
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

/*
The State type represents the state of a URL.
The Pollers send State values to the StateMonitor,
which maintains a map of the current state of each URL.
*/
type State struct {
	url    string
	status string
}

/*
The StateMonitor receives State values on a channel and periodically
outputs the state of all Resources being polled by the program.

maintains a map that stores the state of the URLs being
polled, and prints the current state every updateInterval nanoseconds.
It returns a chan State to which resource state should be sent.
*/
func StateMonitor(updateInterval time.Duration) chan<- State {

	/*
		The variable updatesChannel is a channel of State, on which the Poller goroutines send State values.
		This channel is returned by the function.
	*/
	updatesChannel := make(chan State)

	// The variable urlStatusMap is a map of URLs to their most recent status.
	urlStatusMap := make(map[string]string)

	/*
		A time.Ticker is an object that repeatedly sends a value on a channel at a specified interval.
		In this case, tickerChannel triggers the printing of the current state
		to standard output every updateInterval nanoseconds.
	*/
	tickerChannel := time.NewTicker(updateInterval)

	/*
		Notice that this goroutine owns the urlStatus data structure,
		ensuring that it can only be accessed sequentially.
		This prevents memory corruption issues that might arise from parallel
		reads and/or writes to a shared map.
	*/
	go func() {
		// a for loop which loop forever
		for {
			// The select statement blocks until one of its communications is ready to proceed.
			select {
			case <-tickerChannel.C:
				logState(urlStatusMap)
			case s := <-updatesChannel:
				urlStatusMap[s.url] = s.status
			}
		}
	}()

	return updatesChannel
}

func logState(m1 map[string]string) {
	log.Println("Current state:")
	for k, v := range m1 {
		log.Printf(" %s %s", k, v)
	}
}

/*
A Resource represents the state of a URL to be polled:
the URL itself and the number of errors encountered since the last successful poll.
When the program starts, it allocates one Resource for each URL.
The main goroutine and the Poller goroutines send the Resources to each other on channels.
*/
type Resource struct {
	url      string
	errCount int
}

/*
The Poll method (of the Resource type) performs an HTTP HEAD request
for the Resource's URL and returns the HTTP response's status code.

If an error occurs, Poll logs the message to standard error and returns the error string instead.
*/
func (resource *Resource) Poll() string {
	resp, err := http.Head(resource.url)
	if err != nil {
		log.Println("Error", resource.url, err)
		resource.errCount++
		return err.Error()
	}
	resource.errCount = 0
	return resp.Status
}

/*
Sleep calls time.Sleep to pause before sending the Resource to done.
The pause will either be of a fixed length (pollInterval) plus an additional delay
proportional to the number of sequential errors (r.errCount).

This is an example of a typical Go idiom:
a function intended to run inside a goroutine takes a channel,
upon which it sends its return value (or other indication of completed state).
*/
func (resource *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(resource.errCount))
	done <- resource
}

/*
Each Poller receives Resource pointers from an input channel.
In this program, the convention is that sending a Resource pointer on a channel
passes ownership of the underlying data from the sender to the receiver.

Because of this convention, we know that no two goroutines will access this Resource at the same time.
This means we don't have to worry about locking to prevent concurrent access to these data structures.

The Poller processes the Resource by calling its Poll method.
It sends a State value to the status channel, to inform the StateMonitor of the result of the Poll.

Finally, it sends the Resource pointer to the out channel.
This can be interpreted as the Poller saying "I'm done with this Resource"
and returning ownership of it to the main goroutine.

Several goroutines run Pollers, processing Resources in parallel.
*/
func Poller(input <-chan *Resource, output chan<- *Resource, statusChan chan<- State) {
	for i := range input {
		pollStatus := i.Poll()
		statusChan <- State{i.url, pollStatus}
		output <- i
	}
}

/*
This function starts the Poller and StateMonitor goroutines and
then loops passing completed Resources back to the pending channel after appropriate delays.
*/
func RunUrlPool() {

	/*
		Create our input and output channels:
		First, makes two channels of *Resource, pending and completed.
		a new goroutine sends one Resource per URL to pending and the main goroutine
		receives completed Resources from completed.

		The pending and completed channels are passed to each of the Poller goroutines,
		within which they are known as in and out.
	*/
	pending, completed := make(chan *Resource), make(chan *Resource)

	/*
		StateMonitor will initialize and launch a goroutine that stores the state of each Resource.
		For now, the important thing to note is that it returns a channel of State,
		which is saved as state and passed to the Poller goroutines.
	*/
	state := StateMonitor(statusInterval)

	/*
		launches a number of Poller goroutines, passing the channels as arguments.
		The channels provide the means of communication between
		the main, Poller, and StateMonitor goroutines.
	*/
	for i := 0; i < numPollers; i++ {
		go Poller(pending, completed, state)
	}

	/*
		Send some Resources to the pending queue:
		To add the initial work to the system, we starts a new goroutine that allocates
		and sends one Resource per URL to pending.

		The new goroutine is necessary because unbuffered channel sends and receives are synchronous.

		That means these channel sends will block until the Pollers are ready to read from pending.
		Were these sends performed in the main goroutine with fewer Pollers than channel sends,
		the program would reach a deadlock situation, because main would not yet be receiving from complete.
	*/
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	/*
		When a Poller is done with a Resource, it sends it on the complete channel.
		This loop receives those Resource pointers from complete.

		For each received Resource, it starts a new goroutine calling the Resource's Sleep method.
		Using a new goroutine for each ensures that the sleeps can happen in parallel.

		Note that any single Resource pointer may only be sent on either pending
		or complete at any one time.
		This ensures that a Resource is either being handled by a Poller goroutine or sleeping,
		but never both simultaneously. In this way, we share our Resource data by communicating.
	*/
	for r := range completed {
		go r.Sleep(pending)
	}
}

/*
In this codewalk we have explored a simple example of using Go's concurrency primitives
to share memory through communication.

This should provide a starting point from which to explore the ways
in which goroutines and channels can be used to write expressive and concise concurrent programs.
*/
