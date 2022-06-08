//Package average calculates the average value.
package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/arrays/datafile"
	"log"
)

func main() {
	//numbers := [3]float64{71.8, 56.2, 89.5}
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, numbers := range numbers {
		sum += numbers
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Avarage: %0.2f\n", sum/sampleCount)
}
