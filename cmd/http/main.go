package main

import (
	"fmt"
	"net/http"
)

func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	//	TODO: write the code
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", IndexPathHandler)
	http.ListenAndServe(":3000", nil)
}
