package rotation_test

import (
	"fmt"
	"gotest/rotation"
	"reflect"
	"testing"
)

// Структура тестовых данных, input - входной массив,
// iteration - колличество повторений,
// output - ожидаеммый вывод

type testdata struct {
	input     []int
	iteration int
	output    []int
}

// Входной, повторения, ожидания

var test = []testdata{
	{[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
	{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, -1000}, 1, []int{-1000, 1, 2, 3}},
}

func TestRotation(t *testing.T) {

	t.Run(
		"insert slices", func(t *testing.T) {
			for _, data := range test {
				//	Передаём массив и повторения в функцию и записываем результат в переменную
				v := rotation.Rotation(data.input, data.iteration)
				// Сравниваем вывод с ожиданием
				if reflect.DeepEqual(v, data.output) != true {
					t.Error(
						"For: ", data.input,
						"Expected: ", data.output,
						"Got: ", v,
					)
				} else {
					fmt.Println(
						"For: ", data.input,
						"Iterations", data.iteration,
						"Expected: ", data.output,
						"Got: ", v,
					)

				}
			}
		},
	)

}
