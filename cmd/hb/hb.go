package main

import (
	"flag"

	"github.com/supernetnavi/hummingbird"
)

func main() {
	version := flag.Bool("v", false, "Print hummingbird version")
	port := flag.String("p", "8080", "A port number which hummingbird listens on")
	flag.Parse()

	if *version {
		hummingbird.PrintVersion()
		return
	}

	hummingbird.ServerConfig.Port = *port
	hummingbird.RegisterAPIEndpoints()
	hummingbird.Run()
}
