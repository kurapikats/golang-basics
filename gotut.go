package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

// return multiple
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {

	// auto unpack, type inferred
	// num1, num2 := 5.6, 9.5

	// fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "There"

	fmt.Println(multiple(w1, w2))

}
