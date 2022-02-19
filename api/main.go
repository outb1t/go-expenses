package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		return
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "TEST", r.URL.Path[1:])
	if err != nil {
		return
	}
}
