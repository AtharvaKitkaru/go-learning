package channels

import (
    "fmt"
    "time"
)

// Producer function sends data to the channel
func producer(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        ch <- i
        fmt.Printf("Produced: %d\n", i)
        time.Sleep(500 * time.Millisecond) // Simulate work
    }
    close(ch) // Close channel when done sending
}

// Consumer function receives and processes data from the channel
func consumer(id int, ch <-chan int) {
    for value := range ch {
        fmt.Printf("Consumer %d received: %d\n", id, value)
        time.Sleep(1 * time.Second) // Simulate processing time
    }
}

func CallPC() {
    ch := make(chan int)

    // Start producer goroutine
    go producer(ch)

    // Start multiple consumer goroutines
    for i := 1; i <= 3; i++ {
        go consumer(i, ch)
    }

    // Wait for all goroutines to complete (using a simple sleep here for demonstration purposes)
    time.Sleep(6 * time.Second)
    fmt.Println("Finished processing")
}
