package main

import "fmt"

func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}
	var leters = []string{"a", "b", "c"}
	for i, leter := range leters {
		fmt.Println(i, leter)
	}
}
