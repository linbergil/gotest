package consecutive

func Consecutive(A []int) int {

	if len(A) != 0 {

		min, max := extremum(A)

		if len(A) == max-min+1 && len(A) > 1 {

			return 1

		}

	}

	return 0

}

func extremum(A []int) (int, int) {
	min := A[0]
	max := A[0]

	for i := 0; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}

	}
	return min, max
}
