package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет от микросервиса 2!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
