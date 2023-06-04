package learn

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(message)
		time.Sleep(1 * time.Second)
	}
}

/*
When you run this program, you'll see the messages printed concurrently by the goroutines,
resulting in interleaved output. The exact order of the messages may vary each time you run the program.

This example demonstrates how goroutines can be used to perform tasks concurrently,
allowing multiple functions to execute concurrently and potentially improve the overall execution time of the program.
*/
func RUN_CO1() {
	go printMessage("Hello")   // Start a goroutine for printing "Hello"
	go printMessage("Bonjour") // Start a goroutine for printing "Bonjour"
	go printMessage("Hola")    // Start a goroutine for printing "Hola"

	// Sleep for a while to allow goroutines to execute
	time.Sleep(6 * time.Second)
}
