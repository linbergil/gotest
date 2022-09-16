package unique

func FindUnique(A []int) int {
	var result int

	for i := 0; i < len(A); i++ {
		result = result ^ A[i]
	}

	return result
}
