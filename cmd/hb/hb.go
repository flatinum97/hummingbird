package main

import (
	"flag"

	"github.com/supernetnavi/hummingbird"
)

func main() {
	version := flag.Bool("v", false, "Print hummingbird version")
	flag.Parse()

	if *version {
		hummingbird.PrintVersion()
		return
	}

	hummingbird.RegisterAPIEndpoints()
	hummingbird.Run()
}
