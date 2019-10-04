package hummingbird

import (
	"fmt"
	"net/http"
)

func renderHeaders(w http.ResponseWriter, r *http.Request) {
	for k := range r.Header {
		v := http.Header.Get(r.Header, k)
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
