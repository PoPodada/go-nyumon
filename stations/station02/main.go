package main

import "fmt"

func main() {
	x, y := 1, 2

	fmt.Println("x: ", x, "y: ", y) // 期待される出力: 「x: 1, y: 2」

	Swap(&x, &y)

	fmt.Println("x: ", x, "y: ", y) // 期待される出力: 「x: 2, y: 1」
}

func Swap(x, y *int) {
	*x, *y = *y, *x
}
