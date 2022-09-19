package main

import (
	"fmt"
	"gotest/missing"
)

func main() {

	fmt.Print(missing.FindNumber([]int{1, 3, 5, 2, 6, 7, -9, 8, 4}))

}
