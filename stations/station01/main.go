package main

import "fmt"

func FizzBuzz() {
	for i := 1; i<=100; i++ {
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

func main() {
	fmt.Println("Start FizzBuzz....")

	FizzBuzz() // この関数を実装する

	fmt.Println("FizzBuzz completed!")
}
