// Package main provides

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}

func main() {
	http.HandleFunc("/", myfunc)
	http.ListenAndServe(":8888", nil)
}

//http::localhost:8888/debug/pprof/
