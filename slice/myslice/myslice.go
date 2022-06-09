package main

import "fmt"

func main() {
	//var myArray [5]int
	//var mySlice []string
	//var notes []string
	//notes = make([]string, 7)
	//notes[0] = "do"
	//notes[1] = "re"
	//notes[2] = "mi"
	//fmt.Println(notes[0])
	//fmt.Println(notes[1])

	//primes := make([]int, 5)
	//primes[0] = 2
	//primes[1] = 2

	//notes := make([]string, 7)
	//primes := make([]int, 5)
	//fmt.Print(len(notes))
	//fmt.Print(len(primes))

	//letters := []string{"a", "b", "c"}
	//for i := 0; i < len(letters); i++ {
	//	fmt.Println(letters[i])
	//}
	//for _, letter := range letters {
	//	fmt.Println(letter)
	//}

	//notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	//fmt.Println(notes[3], notes[6], notes[0])
	//primes := []int{
	//	2,
	//	3,
	//	5,
	//}
	//fmt.Println(primes[0], primes[1], primes[2])

	//underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	//i, j := 1, 4
	//slice2 := underlyingArray[i:j]
	//fmt.Println(slice2)
	//slice3 := underlyingArray[2:5]
	//fmt.Println(slice3)

	//slice5 := underlyingArray[1:]
	//fmt.Println(slice5)

	//array1 := [5]string{"a", "b", "c", "d", "e"}
	//slice1 := array1[0:3]
	//array1[1] = "x"
	//fmt.Println(array1)
	//fmt.Println(slice1)

	//array2 := [5]string{"f", "g", "h", "i", "j"}
	//slice2 := array2[2:5]
	//slice2[1] = "X"
	//fmt.Println(array2)
	//fmt.Println(slice2)

	//array3 := [5]string{"a", "b", "c", "d", "e"}
	//slice3 := array3[0:3]
	//slice4 := array3[2:5]
	//array3[2] = "X"
	//fmt.Println(slice3, slice4)

	//slice := []string{"a", "b"}
	//fmt.Println(slice, len(slice))
	//slice = append(slice, "c")
	//fmt.Println(slice, len(slice))

	//slice = append(slice, "d", "e")
	//fmt.Println(slice, len(slice))

	//s1 := []string{"s1", "s1"}
	//s2 := append(s1, "s2", "s2")
	//s3 := append(s2, "s3", "s3")
	//s4 := append(s3, "s4", "s4")
	//fmt.Println(s1, s3, s4)
	//s4[0] = "XX"
	//fmt.Println(s1, s2, s3, s4)
	//

	//floatSlice := make([]float64, 10)
	//boolSlice := make([]bool, 10)
	//fmt.Println(floatSlice[9], boolSlice[5])
	//
	//var intSlice []int
	//var stringSlice []string
	//fmt.Printf("intSlice: %#v, strinsSlice: %#v\n", intSlice, stringSlice)
	//fmt.Println(len(intSlice))
	//intSlice = append(intSlice, 27)
	//fmt.Printf("intSlice: %#v\n", intSlice)
	//
	//var slice []string
	//if len(slice) == 0 {
	//	slice = append(slice, "first item")
	//}
	//fmt.Printf("%#v\n", slice)

	//array := [5]string{"a", "b", "c", "d", "e"}
	//slice := array[1:3]
	//slice = append(slice, "x")
	//slice = append(slice, "y", "z")
	//for _, letter := range slice {
	//	fmt.Println(letter)
	//}
	severalInts(1)
	severalInts(1, 2, 3)
	severalStrings("a", "b")
	severalStrings("a", "b", "c", "d", "e")
	severalStrings()

}
func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
func severalStrings(strings ...string) {
	fmt.Println(strings)
}
