package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(task int) {
	fmt.Printf("Worker is working on task %d\n", task)
	time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1) // Start job increase 1
		i := i

		go func() {
			defer wg.Done() // If this function is done -> decrease 1
			worker(i)
		}()

		wg.Wait() // Wait until zero
	}

}
