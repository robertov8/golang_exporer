package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // taking the address of the variable
	*p++   // dereferenced (taking the value)
	i++

	// Go has no pointer arithmetic
	// p++

	fmt.Println(i, &i, *p, p)
}
