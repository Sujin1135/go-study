package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requested method is %s\n", r.Method)
}

type student struct {
	Name  string
	Age   int
	Score int
}

func newStudent(name string, age, score int) *student {
	return &student{Name: name, Age: age, Score: score}
}

var students []*student

func StudentPostHandler(w http.ResponseWriter, r *http.Request) {
	student := newStudent("철수", 59, 100)
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprint(w, string(data))

	students := append(students, student)
	fmt.Print(students)
}

func NewWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexPathHandler)
	mux.HandleFunc("/students", StudentPostHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}

func main() {
	http.ListenAndServeTLS(":3000", "mycommoncrt.crt", "private.key", NewWebHandler())
}
