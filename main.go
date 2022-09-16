package main

import (
	"fmt"
	"gotest/unique"
)

func main() {
	var nums = []int{2, 2, 4, 5, 4, 6, 6}

	fmt.Println(unique.FindUnique(nums))
}
