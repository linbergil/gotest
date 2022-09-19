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

var odd_test = []testdata{
	{[]int{9, 3, 9, 3, 9, 7, 9}, 7},
	{[]int{9, 3, 9, 8, 9, 3, 9}, 8},
}

var even_test = []testdata{
	{[]int{9, 3, 9, 3, 9, 7}, 0},
}

// сравниваем входные данные с выходными данными, и кидаем ошибки или пишем в терминал
func test(d *testdata, t *testing.T) {

	result := unique.FindUnique(d.input)

	if result != d.output {
		t.Error(
			"test failed",
			"for: ", d.input,
			"expected: ", d.output,
			"got: ", result,
		)
	} else {
		fmt.Println(
			"test passed",
			"for: ", d.input,
			"got: ", result,
		)
	}
}

func TestUnique(t *testing.T) {

	t.Run(
		"inserting odd nums of elements", func(t *testing.T) {

			for _, data := range odd_test {

				test(&data, &testing.T{})

			}
		},
	)

	t.Run(
		"inserting even nums of elements", func(t *testing.T) {

			for _, data := range even_test {

				test(&data, &testing.T{})

			}
		},
	)

}
