package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NumberOfWorkers = 5
)

func processUrl(url string) {
	fmt.Println(fmt.Sprintf("url %v", url))

	// processing
	time.Sleep(5 * time.Second)

	// done
	fmt.Println(fmt.Sprintf("url %v is done", url))
}

func crawl(ch chan string, wg *sync.WaitGroup) {
	for i := 0; i < NumberOfWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range ch {
				processUrl(url)
			}
			return
		}()
	}
}

func main() {
	ch := make(chan string, NumberOfWorkers)
	var wg sync.WaitGroup

	// push the url to channel
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- fmt.Sprintf("%v", i)
		}

		close(ch)
	}()

	// get data from channel and process crawl
	crawl(ch, &wg)

	wg.Wait()
	fmt.Println("Done...")
}
