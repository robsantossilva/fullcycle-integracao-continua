package main

import "fmt"

func main() {
	fmt.Println(sum(2, 2))
}

func sum(a, b int) int {
	result := a + b

	return result
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func sumX(a int, b int) int {
	return a + b + a
}
