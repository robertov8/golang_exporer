package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func printF(value int) {
	fmt.Println(value)
}

func main() {
	result := sum(1, 2)
	printF(result)
}
