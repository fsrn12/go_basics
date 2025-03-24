package lib

import (
	"fmt"
	s "strings"
)

func Exercise() {
	sentence := "the quick brown fox jumps over a lazy dog"
	CapitalizedSentence := Capitalize(sentence)
	fmt.Println(CapitalizedSentence)
	EvenOrOdd()
}
func CapitalizeWord(str string) string {
	return s.ToUpper(s.Split(str, "")[0]) + str[1:]
}
func Capitalize(str string) string {
	strArray := s.Split(str, " ")
	capitalizedSentence := ""
	for _, value := range strArray {
		capitalizedSentence += CapitalizeWord(value) + " "
	}
	return capitalizedSentence
}

func EvenOrOdd() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
