package hummingbird

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func renderHeaders(w http.ResponseWriter, r *http.Request) {
	for k := range r.Header {
		v := http.Header.Get(r.Header, k)
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func renderBody(w http.ResponseWriter, r *http.Request) {
	b := new(bytes.Buffer)
	if _, err := b.ReadFrom(r.Body); err != nil {
		log.Printf("Failed to read HTTP body")
	}
	body := b.String()
	fmt.Fprintf(w, body)
}
}
