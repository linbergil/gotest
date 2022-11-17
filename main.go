package main

import (
	"fmt"
	"gotest/FizzBuzz"
	"gotest/consecutive"
)

func main() {

	array := []int{1, 2, 4, 3, 6, 5}

	fmt.Println("consecutive", consecutive.Consecutive(array))

	FizzBuzz.Fizzbuzz()
}
