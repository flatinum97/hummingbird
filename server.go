package hummingbird

import (
	"log"
	"net/http"
	"strconv"
)

// RegisterAPIEndpoints initializes API Endpoints
func RegisterAPIEndpoints() {
	http.HandleFunc("/headers", renderHeaders)
	http.HandleFunc("/body", renderBody)
}

// Run starts up a HTTP server
func Run() {
	_, err := strconv.Atoi(config.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
