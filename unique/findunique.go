// Чудные вхождения в массив
package unique

import "fmt"

func FindUnique(A []int) int {
	var result int

	if len(A)%2 != 0 {
		for i := 0; i < len(A); i++ {
			//xor use
			result = result ^ A[i]
		}
	} else {
		fmt.Println("the number of elements is even")
	}

	return result
}
