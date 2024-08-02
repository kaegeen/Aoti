package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Aaslema Baba!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}

	result := 0
	ss := 1
	sss := 0

	if s[0] == '-' {
		ss = -1
		sss++
	} else if s[0] == '+' {
		sss++
	}

	for i := sss; i < len(s); i++ {
		char := s[i]
		if char < '0' || char > '9' {
			return 0
		}
		result = result*10 + int(char-'0')
	}

	return result * ss
}
