# requestid [![GoDoc](https://godoc.org/github.com/ascarter/requestid?status.svg)](http://godoc.org/github.com/ascarter/requestid)[![Go Report Card](https://goreportcard.com/badge/github.com/ascarter/requestid)](https://goreportcard.com/report/github.com/ascarter/requestid)

RequestID middleware for Go. RequestID adds a UUID as `X-Request-ID` header if not already set. It also adds it to the http.Request Context. Use `requestid.FromContext` to get the generated request id.

# Example

```go

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ascarter/requestid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rid, _ := requestid.FromContext(r.Context())
	log.Println("Running hello handler:", rid)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	h := http.HandlerFunc(handler)
	http.Handle("/", requestid.RequestIDHandler(h))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

```
