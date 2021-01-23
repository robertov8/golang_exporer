package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	note := 6.9
	noteFinal := int(note)
	fmt.Println(noteFinal)

	// warning: ascii code
	fmt.Println("Test " + string(97))

	// int to string
	fmt.Println("Test ", strconv.Itoa(97))

	// string to int
	num, _ := strconv.Atoi("97")
	fmt.Println("Test ", num-95)

	// parse bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("True")
	}
}
