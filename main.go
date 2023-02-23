package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "ContentLength: %d\n", r.ContentLength)
	fmt.Fprintf(w, "TransferEncoding: %s\n", r.TransferEncoding)
	fmt.Fprintf(w, "Close: %t\n", r.Close)
	fmt.Fprintf(w, "Trailer: %s\n", r.Trailer)
	fmt.Fprintf(w, "User Agent: %s\n", r.UserAgent())
	fmt.Fprintf(w, "Remote Address: %s\n", r.RemoteAddr)

	fmt.Fprintln(w, "\nHeaders:")

	// Sort the headers for a consistent output
	var keys []string
	for key := range r.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		values := r.Header[key]
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	}

}

func main() {
	http.HandleFunc("/", headersHandler)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
