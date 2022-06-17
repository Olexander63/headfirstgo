package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println("Sleepy Creek Country populaton:", population)
	fmt.Println("Coongratulations, Kevin and Anna! it's a girl!")
	population += 1
	fmt.Println("Sleepy Creek population:", population)
}
