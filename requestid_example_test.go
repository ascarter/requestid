package requestid_test

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/ascarter/requestid"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rid, ok := ctx.Value(ridKey).(string)
	log.Println("Running hello handler")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Example() {
	h := http.HandlerFunc(handler)
	http.Handle("/", requestid.RequestIDHandler(h))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
