package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stop := make(chan bool)
	crashed := make(chan string) // Channel to receive crash info

	wg.Add(2)

	// Start goroutine 1
	go routine1(stop, &wg, crashed)

	// Start goroutine 2
	go routine2(stop, &wg, crashed)

	// Wait for crash signal or timeout
	select {
	case crashMsg := <-crashed:
		// One goroutine crashed - stop all
		fmt.Println("ALERT:", crashMsg)
		close(stop) // Close channel to signal all goroutines

	case <-time.After(10 * time.Second):
		// Safety timeout
		fmt.Println("Timeout - stopping all")
		close(stop)
	}

	wg.Wait()
	fmt.Println("Main: All goroutines stopped")
}

func routine1(stop chan bool, wg *sync.WaitGroup, crashed chan string) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		select {
		case <-stop:
			fmt.Println("Routine 1: Received stop signal, exiting...")
			return
		default:
			// Simulate crash at iteration 3
			if i == 3 {
				crashed <- "Routine 1 crashed at iteration 3!" // Send crash signal
				return
			}
			fmt.Println("Routine 1: Working...", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func routine2(stop chan bool, wg *sync.WaitGroup, crashed chan string) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		select {
		case <-stop:
			fmt.Println("Routine 2: Received stop signal, exiting...")
			return
		default:
			fmt.Println("Routine 2: Working...", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
