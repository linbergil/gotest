package consecutive_test

import (
	"fmt"
	"gotest/consecutive"
	"testing"
)

// Структура тестовых данных, input - входной массив,
// output - ожидаеммый вывод

type testdata struct {
	input  []int
	output int
}

// Входной, повторения, ожидания

var test = []testdata{
	{[]int{1, 2, 4, 3, 6, 5}, 1},
	{[]int{1, 2, 3, 5}, 0},
	{[]int{255, 256, 257, 258}, 1},
}

func TestCons(t *testing.T) {

	t.Run(
		"insert slices", func(t *testing.T) {
			for _, data := range test {
				//	Передаём массив и повторения в функцию и записываем результат в переменную
				v := consecutive.Consecutive(data.input)
				// Сравниваем вывод с ожиданием
				if v != data.output {
					t.Error(
						"For: ", data.input,
						"Expected: ", data.output,
						"Got: ", v,
					)
				} else {
					fmt.Println(
						"For: ", data.input,
						"Expected: ", data.output,
						"Got: ", v,
					)

				}
			}
		},
	)

}
