package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("https://www.github.com")
	if err != nil {
		panic(err)
	}
	// without body
	nobody, err := httputil.DumpResponse(resp, false)
	if err != nil {
		panic(err)
	}
	// with body
	withbody, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp line and header len: %d, body len: %d\n", len(nobody), len(withbody))
}
