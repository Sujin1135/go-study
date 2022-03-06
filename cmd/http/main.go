package main

import (
	"fmt"
	"net/http"
)

func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested method is %s\n", r.Method)
}

func NewWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexPathHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}

func main() {
	http.ListenAndServe(":3000", NewWebHandler())
}
