package solutions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotest/consecutive"
	"gotest/missing"
	"gotest/rotation"
	"gotest/unique"
	"io"
	"log"
	"net/http"
	"sync"
)

type taskElement struct {
	a      []int
	k      int
	result []int
}

func ProcessTask(taskName string) ([]byte, error) {
	//получаем условия для задач
	var taskCases []json.RawMessage
	err := getCases(taskName, &taskCases)
	if err != nil {
		log.Fatalln(err)
	}
	//решение задач
	taskArray := [10]taskElement{}
	err = solve(taskName, taskCases, &taskArray)
	if err != nil {
		log.Fatalln(err)
	}
	//сборка результата
	var data []byte
	data, err = requestReview(taskName, &taskCases, &taskArray)
	if err != nil {
		log.Fatalln(err)
	}

	return data, nil
}

func getCases(taskName string, taskCases *[]json.RawMessage) error {

	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", SolutionURL, taskName)) //протягиваем руку за подоянием
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	payload, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(payload, &taskCases)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
func solve(taskName string, taskCases []json.RawMessage, taskArray *[10]taskElement) error {
	var wg sync.WaitGroup //ждёмс рутины
	wg.Add(10)

	for i, taskCase := range taskCases {
		//парсим
		var arguments []json.RawMessage
		err := json.Unmarshal(taskCase, &arguments)
		if err != nil {
			log.Fatalln(err)
		}

		if taskName == Rotation { //ротация отличается от других массивов
			json.Unmarshal(arguments[0], &taskArray[i].a)
			json.Unmarshal(arguments[1], &taskArray[i].k)
		} else {
			err := json.Unmarshal(arguments[0], &taskArray[i].a)
			if err != nil {
				log.Fatalln(err)
				return err
			}
		}
		//отправляем входные данные в задачки
		go func(i int) {
			defer wg.Done() //дождались
			switch taskName {
			case Rotation:
				taskArray[i].result = rotation.Rotation(taskArray[i].a, taskArray[i].k)
			case FindUnique:
				taskArray[i].result = append(taskArray[i].result, unique.FindUnique(taskArray[i].a))
			case Cons:
				taskArray[i].result = append(taskArray[i].result, consecutive.Consecutive(taskArray[i].a))
			case Missing:
				taskArray[i].result = append(taskArray[i].result, missing.FindNumber(taskArray[i].a))
			default:
			}
		}(i)

	}

	wg.Wait() //ждёмс
	return nil
}

func requestReview(taskName string, taskCases *[]json.RawMessage, taskArray *[10]taskElement) ([]byte, error) {
	//собраем решения задач
	var result []byte
	var results []json.RawMessage
	var err error

	for _, element := range taskArray {
		if taskName == Rotation {
			result, err = json.Marshal(element.result)
		} else {
			result, err = json.Marshal(element.result[0])
		}
		if err != nil {
			log.Fatalln(err)
		}
		results = append(results, result)

	}

	message := map[string]interface{}{
		"user_name": userName,
		"task":      taskName,
		"results": map[string]interface{}{
			"payload": taskCases,
			"results": results,
		},
	}

	packedMessage, err := json.Marshal(message)
	if err != nil {
		log.Fatalln()
	}

	request, err := http.Post(fmt.Sprintf("%s/tasks/solution", SolutionURL), "application/json", bytes.NewBuffer(packedMessage))

	defer request.Body.Close()

	var data []byte

	if request.StatusCode != http.StatusOK {
		data = []byte(request.Status)
	} else {
		data, err = io.ReadAll(request.Body)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return data, nil
}
