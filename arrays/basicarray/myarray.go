package main

import "fmt"

func main() {
	// array strings
	//var notes [7]string
	//notes[0] = "do"
	//notes[1] = "re"
	//notes[2] = "mi"
	//fmt.Println(notes[0])
	//fmt.Println(notes[1])

	//array int number
	//var primes [5]int
	//primes[0] = 2
	//primes[1] = 3
	//fmt.Println(primes[0])

	//array time

	//var dates [3]time.Time
	//dates[0] = time.Unix(1257894000, 0)
	//dates[1] = time.Unix(1447920000, 0)
	//dates[2] = time.Unix(1508632200, 0)
	//fmt.Println(dates[1])

	//counters
	//var counters [3]int
	//counters[0]++
	//counters[0]++
	//counters[2]++
	//fmt.Println(counters[0], counters[0], counters[2])

	//var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	//notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	//index := 1
	//fmt.Println(index, notes[index])
	//index = 3
	//fmt.Println(index, notes[index])
	//fmt.Println(notes[3], notes[6], notes[0])
	//var primes [5]int = [5]int{2, 3, 5, 7, 11}
	//primes := [5]int{2, 3, 5, 7, 11}
	//fmt.Println(primes[0], primes[2], primes[4])
	//fmt.Println(notes)
	//fmt.Println(primes)
	//fmt.Println("---------------")
	//fmt.Printf("%#v\n", notes)
	//fmt.Printf("%#v\n", primes)

	//text := [3]string{
	//	"This is a series of long strings",
	//	"which would be awkward to place",
	//	"together on a single line",
	//}
	//fmt.Println(text[2])
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(len(notes))
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(len(primes))
	for index, note := range notes {
		fmt.Println(index, note)
	}
}
