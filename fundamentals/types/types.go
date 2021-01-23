package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numbers int
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal number is", reflect.TypeOf(32999))

	// numbers unsigned = uint8 uint16 uint16 uint64
	var b byte = 255
	fmt.Println("The byte is", reflect.TypeOf(b))

	// numbers signed = int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("The max value is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	// representation table unicode (int32)
	var i2 rune = 'a'
	fmt.Println("The rune is", reflect.TypeOf(i2))
	fmt.Println("The value rune is", i2)

	// real numbers (float32, float64)
	var x float32 = 49.99
	fmt.Println("The type of x is", reflect.TypeOf(x))
}
