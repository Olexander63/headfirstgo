package main

import (
	"fmt"
	"log"
)

// global
//var metersPerLiter float64

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is valid", height)
	}
	area := width * height
	return area / 10.0, nil
}
func main() {
	//var amount, total float64
	//amount = paintNeeded(4.2, 3.0)
	//fmt.Printf("%.2f liters needed\n", amount)
	//total += amount
	//amount = paintNeeded(5.2, 3.5)
	//fmt.Printf("%0.2f liters needed\n", amount)
	//total += amount
	//fmt.Printf("Total: %0.2f liters\n", total)

	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	//after added func paintNeeded
	//(4.2, 3.0)
	//paintNeeded(5.2, 3.5)
	//paintNeeded(5.0, 3.3)

	//var width, height, area float64
	//
	//width = 4.2
	//height = 3.0
	//area = width * height
	//fmt.Printf("%.2f liters needs\n", area/10.0)
	//width = 5.2
	//height = 3.5
	//area = width * height
	//fmt.Printf("%.2f liters needs\n", area/10.0)

}
