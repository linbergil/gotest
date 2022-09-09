package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz\n")
		} else if i%3 == 0 {
			fmt.Print("Fizz\n")
		} else if i%5 == 0 {
			fmt.Print("Buzz\n")
		} else {
			fmt.Println(i)
		}
	}

	switchFB()
}

func switchFB() {
	for i := 1; i < 101; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

}
