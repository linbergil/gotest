package rotation

func Rotation(A []int, K int) []int {

	for i := 0; i < K; i++ {
		//A[len(A)-1:] - последний элемент; A[:len(A)-1] - все элементы до последнего
		A = append(A[len(A)-1:], A[:len(A)-1]...)

	}

	return A
}
