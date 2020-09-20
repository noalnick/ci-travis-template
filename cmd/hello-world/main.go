package main

import "fmt"
import math "github.com/ctolon22/ci-travis/internal/math"

func main() {
	fmt.Println("SUM: 2 + 3")
	fmt.Println(math.Sum(2, 3))

	fmt.Println("SUBSTRACT: 10 - 8")
	fmt.Println(math.Substract(10, 8))

	fmt.Println("SQUARE: 3")
	fmt.Println(math.Square(3))

	fmt.Println("Mult: 3 * 5")
	fmt.Println(math.Mult(3, 5))
}