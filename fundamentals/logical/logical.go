package main

import "fmt"

func purchase(work1, work2 bool) (bool, bool, bool) {
	purchaseTv50 := work1 && work2
	purchaseTv32 := work1 != work2 // XOR
	purchaseIceCream := work1 || work2

	return purchaseTv50, purchaseTv32, purchaseIceCream
}

func main() {
	tv50, tv32, iceCream := purchase(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, iceCream: %t, Healthy: %t", tv50, tv32, iceCream, !iceCream)
}
