package main

import "fmt"

// No ternary operator
func getResult(note float64) string {
	if note >= 6 {
		return "Approve"
	}

	return "Disapproved"
}

func main() {
	fmt.Println(getResult(6.2))
}
