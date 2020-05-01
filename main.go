package main

import (
	"flag"
	"fmt"
	"httphash/core"
	"os"
)

func main() {

	//number of parallel requests
	parallel := flag.Int("parallel", 10, "number pf parallel requests")
	//fmt.Print(*parallel)
	//fmt.Print("\n")

	arg := os.Args[2:]

	Launcher(arg, parallel)
}

func Launcher(urls []string, parallel *int) {

	for i := 1; i < len(urls); i++ {
		
		validUrl, requestResponse, isValid := core.RequestUrl(urls[i])

		if !isValid {
			fmt.Print("Url format is wrong")
		} else {
			hash := core.HashRequestResponse(requestResponse)

			fmt.Print(validUrl + " " + " " + hash)
			fmt.Print("\n")
		}
	}

}
