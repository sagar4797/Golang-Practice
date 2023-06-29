package main

import "fmt"

func main() {
	fmt.Println("this is string")
	str := "Sagar Patil"
	fmt.Println("Str:", str)
	var result string
	for _, v := range str {
		char := string(v)
		fmt.Println("char:", char)
		result = string(v) + result
	}

	fmt.Println("result:", result)
}
