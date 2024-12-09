package main

import (
	"fmt"
	"twg/math"
)

func main() {
	sum := math.Sum([]int{10, 2, 3})
	if sum != 15 {
		msg := fmt.Sprintf("FAIL: Wanted 15, but got %d", sum)
		panic(msg)
	}
	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: Wanted 15, but got %d", sum)
		panic(msg)
	}
	fmt.Println("PASS")
}
