package main

import (
	"fmt"
	"gotest/solutions"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Engage!")
	testHandler := func(w http.ResponseWriter, _ *http.Request) {
		solutions.ProcessTask(solutions.FindUnique)
		io.WriteString(w, "я стану хокаге!!!")
	}

	http.HandleFunc("/status", testHandler)

	err := http.ListenAndServe(solutions.Port, nil)
	if err != nil {
		panic(err)
	}
}
