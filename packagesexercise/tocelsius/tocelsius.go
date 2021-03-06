package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/packagesexercise/keyboard"
	"log"
)

// tocelsius converts temperature to Fahrenheit
// to degrees Celsius.

//func getFloat() (float64, error) {
//	reader := bufio.NewReader(os.Stdin)
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		return 0, err
//	}
//
//	input = strings.TrimSpace(input)
//	number, err := strconv.ParseFloat(input, 64)
//	if err != nil {
//		return 0, err
//	}
//	return number, nil
//}

func main() {
	fmt.Print("Enter a temperature in Fahrenheit: ")
	fahrengeit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrengeit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Clsius\n", celsius)
}
