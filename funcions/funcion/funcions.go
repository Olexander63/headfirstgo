package main

import "fmt"

func sayHi() {
	fmt.Println("Hi !")
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
func main() {
	sayHi()
	repeatLine("hello", 3)
}
