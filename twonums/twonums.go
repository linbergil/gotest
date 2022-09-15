package twonums

import "fmt"

func main() {
	var test_nums []int
	test_nums = append(test_nums, 2, 3, 5, 6, 9, 1)
	fmt.Print(twoSum(test_nums, 10))
}

func twoSum(nums []int, target int) []int {

	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}
