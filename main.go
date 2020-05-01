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

		//fmt.Print(arg[i])
		//fmt.Print("\n")
		/*resp, _ := http.Get(arg[i])
		fmt.Print(resp)*/

		requestResponse, isUrlReachable := core.RequestUrl(urls[i])

		if !isUrlReachable {
			fmt.Print("Url format is wrong")
		} else {
			hash := core.HashRequestResponse(requestResponse)

			fmt.Print(urls[i] + " " + " " + hash)
			fmt.Print("\n")
		}
	}

}
