package main

import (
	"fmt"
)

func main() {
	s := "suvarna chavan"
	str := []byte(s)
	lenght := len(s) - 1
	for i, j := 0, lenght; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	fmt.Println("THis is str : ", str)
	fmt.Println("THis is str : ", string(str))
}
