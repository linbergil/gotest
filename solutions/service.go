package solutions

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ProcessTask(taskName string) ([]byte, error) {
	//получаем условия для задач
	var taskCases []json.RawMessage
	err := getCases(taskName, taskCases)
	if err != nil {
		log.Fatalln(err)
		return []byte{}, err
	}
	//решение задач
	//

	//сборка результата
	//

	//должен вернуть вывод решения
	return nil, nil
}

func getCases(taskName string, taskCases []json.RawMessage) error {

	response, err := http.Get("http://" + SolutionURL + "/tasks/" + taskName)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(response.Body)

	payload, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = json.Unmarshal(payload, &taskCases)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
