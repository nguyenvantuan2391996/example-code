package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

var g singleflight.Group

func lightweightOperation(param string) (interface{}, error) {
	fmt.Println("Executing function with param:", param)
	return "Result from " + param, nil
}

func main() {
	var wg sync.WaitGroup

	// Simulate goroutines calling the same function concurrently
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Using single flight to deduplicate
			result, err, _ := g.Do("key1", func() (interface{}, error) {
				return lightweightOperation("key1")
			})

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Printf("Goroutine %d received result: %s\n", i, result)
		}(i)
	}

	wg.Wait()
}
