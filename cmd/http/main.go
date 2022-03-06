package main

import (
	"fmt"
	"net/http"
)

func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested method is %s\n", r.Method)
}

func main() {
	http.HandleFunc("/", IndexPathHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
