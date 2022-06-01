package main

import "fmt"

func functionA(a int, b int) {
	fmt.Println(a + b)
}
func functionB(a int, b int) {
	fmt.Println(a * b)
}

func functionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}

func double(number float64) float64 {
	return number * 2
}
func main() {

	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))

	functionA(2, 3)
	functionB(2, 3)
	functionC(true)
	functionD("S", 4)
	functionA(5, 6)
	functionB(5, 6)
	functionC(false)
	functionD("ha", 3)
}
