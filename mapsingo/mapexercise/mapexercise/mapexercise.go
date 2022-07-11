package main

import "fmt"

func main() {
	//var ranks map[string]int
	//ranks = make(map[string]int)

	//ranks := make(map[string]int)

	//ranks["gold"] = 1
	//ranks["silver"] = 2
	//ranks["bronze"] = 3
	//fmt.Println(ranks["bronze"])
	//fmt.Println(ranks["gold"])

	//elements := make(map[string]string)
	//elements["H"] = "hydrogen"
	//elements["Li"] = "Lithium"
	//fmt.Println(elements["Li"])
	//fmt.Println(elements["H"])

	//isPrime := make(map[int]bool)
	//isPrime[4] = false
	//isPrime[7] = true
	//fmt.Println(isPrime[4])
	//fmt.Println(isPrime[7])

	//ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	//fmt.Println(ranks["gold"])
	//fmt.Println(ranks["bronze"])
	//elements := map[string]string{
	//	"H":  "hydrogen",
	//	"Li": "Lithium",
	//}
	//fmt.Println(elements["H"])
	//fmt.Println(elements["Li"])

	//numbers := make(map[string]int)
	//numbers["I've ben assigned"] = 12
	//fmt.Printf("%#v\n", numbers["I've been assigned"])
	//fmt.Printf("%#v\n", numbers["I've haven't been assigned"])

	//counters := make(map[string]int)
	//counters["a"]++
	//counters["a"]++
	//counters["c"]++
	//fmt.Println(counters["a"], counters["b"], counters["c"])

	//var myMap map[int]string = make(map[int]string)
	//myMap[3] = "three"
	//fmt.Printf("%#v\n", myMap)

	//counters := map[string]int{"a": 3, "b": 0}
	//var value int
	//var ok bool
	//value, ok = counters["a"]
	//fmt.Println(value, ok)
	//_, ok = counters["b"]
	//fmt.Println(value, ok)
	//_, ok = counters["c"]
	//fmt.Println(value, ok)

	//data := []string{"a", "c", "e", "a", "e"}
	//counts := make(map[string]int)
	//for _, item := range data {
	//	counts[item]++
	//}
	//letters := []string{"a", "b", "c", "d", "e"}
	//for _, letter := range letters {
	//	counts, ok := counts[letter]
	//	if !ok {
	//		fmt.Printf("%s: not found\n", letter)
	//	} else {
	//		fmt.Printf("%s: %d\n", letter, counts)
	//	}
	//}
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("primeL %v, ok: %v\n", prime, ok)
}
