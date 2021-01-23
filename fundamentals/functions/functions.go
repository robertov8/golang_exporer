package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func print_f(value int) {
	fmt.Println(value)
}

func main() {
	result := sum(1, 2)
	print_f(result)
}
