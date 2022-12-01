package main

import (
	"bytes"
	"fmt"
	"gotest/solutions"
	"log"
	"net/http"
	"strings"
	"sync"
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

	if r.URL.Path == "/tasks/" {
		var answers = make([][]byte, 4)
		var err error
		var wg sync.WaitGroup //ждёмс рутины
		wg.Add(4)

		go func() {
			defer wg.Done()
			answers[0], err = solutions.ProcessTask(solutions.Rotation)
			if err != nil {
				log.Fatalln(err)
			}
		}()

		go func() {
			defer wg.Done()
			answers[1], err = solutions.ProcessTask(solutions.FindUnique)
			if err != nil {
				log.Fatalln(err)
			}
		}()

		go func() {
			defer wg.Done()
			answers[2], err = solutions.ProcessTask(solutions.Cons)
			if err != nil {
				log.Fatalln(err)
			}
		}()

		go func() {
			defer wg.Done()
			answers[3], err = solutions.ProcessTask(solutions.Missing)
			if err != nil {
				log.Fatalln(err)
			}
		}()

		wg.Wait()
		w.Write(bytes.Join(answers, []byte{}))

	} else {

		taskName := strings.Split(r.URL.Path, "/")
		//fmt.Println(taskName[2])
		answer, err := solutions.ProcessTask(taskName[2])
		if err != nil {
			log.Fatalln(err)
		}

		w.Write(answer)

	}

}
