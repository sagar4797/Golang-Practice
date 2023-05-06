package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Println("Enter the number to convert into roman:")
	fmt.Scanf("%d", &n)
	// intToRomanUsingSimpleUsingForLoop(n)
	intToRomanByDivisibility(n)
}

// using substraction method
func intToRomanUsingSimpleUsingForLoop(n int) {

	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "D", "CM", "M"}

	index := len(numbers) - 1
	var roman string

	for n > 0 {
		for numbers[index] <= n {
			roman = roman + romans[index]
			n -= numbers[index]
		}
		index -= 1
	}
	fmt.Printf("\n Using substraction =>Roman is:%s", roman)
}

// using divisibility method
func intToRomanByDivisibility(n int) {

	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 500, 900, 1000}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "D", "CM", "M"}

	index := len(numbers) - 1
	var roman string

	for n > 0 {
		if (n / numbers[index]) > 0 {
			roman = roman + strings.Repeat(romans[index], (n/numbers[index]))
			n = n % numbers[index]
		}
		index -= 1
	}
	fmt.Println("\n Using divisibility Roman is:", roman)
}
