package main

import "fmt"

func main() {
	fmt.Print("same")
	fmt.Print(" line.")

	fmt.Println(" New")
	fmt.Println("line.")

	x := 3.141516
	xs := fmt.Sprint(x)

	fmt.Println("The value x is", x)
	fmt.Println("The value x is " + xs)
	fmt.Printf("The value x is %.2f", x)

	const (
		a = 1
		b = 1.9999
		c = false
		d = "opa"
	)

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v %v", a, b, b, c, d)
}
