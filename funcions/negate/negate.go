package main

import "fmt"

// test
func negate(myboolean bool) bool {
	return !myboolean
}

func main() {
	truth := true
	negate(truth)
	fmt.Println(truth)
	lies := false
	negate(lies)
	fmt.Println(lies)
}
