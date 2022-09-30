package missing

import "fmt"

func FindNumber(A []int) int {

	var elements_sum int

	for i := 0; i < len(A); i++ {

		if A[i] < 0 || A[i] > 100000 {
			fmt.Print("array has negative element or longer than expected \n")
			return A[i]
		} else {
			elements_sum = elements_sum + A[i]
		}

	}

	nums_sum := (len(A) + 1) * (len(A) + 2) / 2

	return nums_sum - elements_sum

}
