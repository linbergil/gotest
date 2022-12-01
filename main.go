package main

import (
	"fmt"
	"gotest/solutions"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/task/", Handler)
	http.HandleFunc("/tasks/", Handler)
	fmt.Println("Engage!")
	err := http.ListenAndServe(solutions.Port, nil)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "я стану хокаге!!!")
}
