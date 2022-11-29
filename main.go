package main

import (
	"fmt"
	"gotest/solutions"
	"net/http"
)

func main() {
	fmt.Println("Live long and prosper")
	testHandler := func(w http.ResponseWriter, _ *http.Request) {
		b, _ := solutions.ProcessTask(solutions.Rotation) //воткнул что бы не ругался на не использование
		//io.WriteString(w, string(b))
		fmt.Println(b)
	}

	http.HandleFunc("/status", testHandler)

	err := http.ListenAndServe(solutions.Port, nil)
	if err != nil {
		panic(err)
	}
}
