package hummingbird

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
)

func renderFullyRequestDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[Method]\n")
	renderRequestMethod(w, r)
	fmt.Fprintf(w, "[Header]\n")
	renderHeaders(w, r)
	fmt.Fprintf(w, "[Body]\n")
	renderBody(w, r)
	fmt.Fprintf(w, "[Remote address]\n")
	renderRemoteAddr(w, r)
}

func renderRequestMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.Method)
}

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

func renderRemoteAddr(w http.ResponseWriter, r *http.Request) {
	remote_ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Remote addr %s is invalid format. Expect <IP>:<PORT>\n", r.RemoteAddr)
	}
	fmt.Fprintf(w, remote_ip)
}
