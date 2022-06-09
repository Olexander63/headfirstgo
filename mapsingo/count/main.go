// Package count counts the number of occurrences
// each line in the file.
package main

import (
	"fmt"
	"github.com/Olexander63/headfirstgo/arrays/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
