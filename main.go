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

	// make sure the app doesn't finish prematurely until all goroutines finish executing
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(urls))


	sem := make(chan struct{}, *numberParallelRequests)

	for _, url := range urls {

		sem <- struct{}{}

		go func(url string) {

			defer func() {
				waitGroup.Done()
				<-sem
			}()

			sendRequest(url)
		}(url)
	}
	// wait until all goroutine finish executing
	waitGroup.Wait()
}

func sendRequest(url string) {
	validUrl, requestResponse, isValid := core.RequestUrl(url)
	if !isValid {
		fmt.Print(validUrl + " Url format is wrong \n")
	} else {
		hash := core.HashRequestResponse(requestResponse)

		fmt.Print(validUrl + " " + " " + hash)
		fmt.Print("\n")
	}
}
