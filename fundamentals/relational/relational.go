package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 > 2)
	fmt.Println(">=", 3 > 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Dates:", d1 == d2)
	fmt.Println("Dates:", d1.Equal(d2))

	type People struct {
		Name string
	}

	p1 := People{"João"}
	p2 := People{"João"}
	fmt.Println("People:", p1 == p2)
}
