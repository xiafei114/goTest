package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go func(url string) {
			// Launch a goroutine to fetch the URL.
			defer wg.Done()
			// Fetch the URL.
			fmt.Println(url)
		}(url)
	}
	// Wait for all goroutines to finish.
	wg.Wait()
	fmt.Println("Game Over")
}
