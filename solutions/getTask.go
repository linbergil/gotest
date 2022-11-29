package solutions

import (
	"io"
	"log"
	"net/http"
)

func NewService() taskService {
	return taskService{}
}

type taskService struct {
	store solutionResult
}

func (t taskService) MakeGet(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	//b, _ := asn1.MarshalWithParams()
	//t.store.Add(Payload(b))
	//
	//fmt.Println(b)
}

//func FindUnique() {
//	fmt.Println()
//
//	//for i, _ := range c {
//	//	fmt.Println(i)
//	//}
//
//}
