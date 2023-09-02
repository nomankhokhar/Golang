// package main

// import (
// 	"fmt"
// 	"time"
// )

// func heavy() {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("Heavy")
// 	}
// }

// func superHeavy() {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		fmt.Println("Super Heavy")
// 	}
// }

// func main() {
// 	fmt.Println("GoROutine")
// 	go heavy()
// 	go superHeavy()
// 	fmt.Println("Fin")
// select {}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	Sum1 := make(chan int)
// 	Sum2 := make(chan int)

// 	go func() {
// 		val := 0
// 		for i := 0; i < 10; i++ {
// 			val += i
// 		}
// 		Sum2 <- val
// 	}()

// 	go func() {
// 		val := 0
// 		for i := 0; i < 10; i++ {
// 			val += i
// 		}
// 		Sum1 <- val
// 	}()

// 	select {
// 	case value := <-Sum1:
// 		fmt.Println("Received from One:", value)
// 	case value2 := <-Sum2:
// 		fmt.Println("Received from Two:", value2)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	ch3 := make(chan int)

// 	go func() {
// 		ch1 <- 42
// 	}()

// 	go func() {
// 		ch2 <- 84
// 	}()

// 	go func() {
// 		ch3 <- 123
// 	}()

// 	select {
// 	case value1 := <-ch1:
// 		fmt.Println("Received from ch1:", value1)
// 	case value2 := <-ch2:
// 		fmt.Println("Received from ch2:", value2)
// 	case value3 := <-ch3:
// 		fmt.Println("Received from ch3:", value3)
// 	}
// }

// This program will received all value of goroutine

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg.Add(3)

	go func() {
		defer wg.Done()
		ch1 <- 42
		close(ch1)

	}()

	go func() {
		defer wg.Done()
		ch2 <- 84
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		ch3 <- 123
		close(ch3)
	}()

	for {
		select {
		case value1, ok := <-ch1:
			if !ok {
				ch1 = nil
			}
			if ok {
				fmt.Println("Received from ch1:", value1)
			}
		case value2, ok := <-ch2:
			if !ok {
				ch2 = nil
			}
			if ok {
				fmt.Println("Received from ch2:", value2)
			}
		case value3, ok := <-ch3:
			if !ok {
				ch3 = nil
			}
			if ok {
				fmt.Println("Received from ch3:", value3)
			}
		}

		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}
	}

	wg.Wait()
	fmt.Println("All values received, terminating the program.")
}
