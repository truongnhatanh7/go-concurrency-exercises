package main

import (
	"fmt"
	"net/http"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is a greet handler")
	w.Write([]byte("Hello, World!"))
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this is a time handler")
	w.Write([]byte("The time is 12:00"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", GreetHandler)
	mux.HandleFunc("/greet/deep", TimeHandler)

	http.ListenAndServe(":8080", mux)
}
