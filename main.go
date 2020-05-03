package main

import (
	"flag"
	"httphash/core"
	"os"
)

func main() {
	var urls []string

	parallel := flag.Int("parallel", 10, "number of parallel requests")
	flag.Parse()

	flagSet := core.IsFlagPassed("parallel")

	if flagSet {
		urls = os.Args[2:]
	} else {
		urls = os.Args[1:]
	}

	//fmt.Println("Flag is set: ", flagSet)
	//fmt.Println("Number of parallel requests: ", *parallel)


	core.Execute(urls, parallel)
}
