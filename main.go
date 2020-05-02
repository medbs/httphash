package main

import (
	"flag"
	"fmt"
	"httphash/core"
	"os"
	"sync"
)

func main() {

	parallel := flag.Int("parallel", 10, "number pf parallel requests")

	arg := os.Args[2:]

	Launcher(arg, parallel)
}

func Launcher(urls []string, numberParallelRequests *int) {

	// make sure the app doesn't finish prematurely until all goroutines finish executing by acting as a counter
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(urls))

	// create a channel that used to make sure that only N=numberParallelRequests goroutines will run at a time
	// use channel of empty struct instead of other types because empty structs occupies zero bytes of storage

	sem := make(chan struct{}, *numberParallelRequests)

	for _, url := range urls {

		//put an empty struct in channel sem. if sem is full, the process won't continue to the next line
		sem <- struct{}{}

		go func(url string) {

			// when this goroutine finishes, make sure to mark the waitGroup for this goroutine as finished;
			//and free the sem variable, so the next goroutine can be executed.
			defer func() {
				waitGroup.Done()
				<-sem
			}()

			//send request to a given url
			sendRequest(url)
		}(url)
	}
	// wait until all goroutine finish executing
	waitGroup.Wait()
}

func sendRequest(url string) {
	validUrl, requestResponse, isValid := core.RequestUrl(url)
	if !isValid {
		fmt.Print(validUrl + " Url format is wrong or cannot be requested \n")
	} else {
		hash := core.HashRequestResponse(requestResponse)

		fmt.Print(validUrl + " " + " " + hash)
		fmt.Print("\n")
	}
}
