package unique_test

import (
	"fmt"
	"gotest/unique"
	"testing"
)

// Структура тестовых данных, input - входной массив,
// output - ожидаеммый вывод

type testdata struct {
	input  []int
	output int
}

var test = []testdata{
	{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	{[]int{9, 3, 9, 3, 9, 7, 9}, 8},
}

func TestRotation(t *testing.T) {

	t.Run(
		"insert nums", func(t *testing.T) {

			for _, data := range test {

				result := unique.FindUnique(data.input)

				if result != data.output {
					t.Error(
						"For: ", data.input,
						"Expected: ", data.output,
						"Got: ", result,
					)
				} else {
					fmt.Println(
						"For: ", data.input,
						"Got: ", result,
					)
				}
			}
		},
	)

}
