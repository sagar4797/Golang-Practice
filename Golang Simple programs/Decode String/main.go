package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	decodedString := decodeString("1000000[leetcode]")
	fmt.Println("This is decodedString: %s", decodedString)
}

// I solved this program using stack
func decodeString(s string) string {
	strStack := []string{} // This stack is stored all char and open bracket('[')
	numStack := []int{}    // This stack only stored numbers
	FinalDecodedString := ""
	numTag := false
	finalNum := 0

	for _, char := range s {
		if char >= '0' && char <= '9' {

			numStr := string(char)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Errorf("Error : %+v", err)
				return ""
			}
			if numTag {
				finalNum = finalNum*10 + num
				numStack[len(numStack)-1] = finalNum

			} else {
				numStack = append(numStack, num)
				finalNum = num
			}
			numTag = true

		} else if char == ']' {
			numTag = false
			finalNum = 0
			decodeStr := ""
			for len(strStack) > 0 && strStack[len(strStack)-1] != "[" {
				decodeStr = strStack[len(strStack)-1] + decodeStr
				strStack = strStack[:len(strStack)-1]
			}

			strStack = strStack[:len(strStack)-1]
			if len(numStack) > 0 {
				currNum := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				decodeStr = strings.Repeat(decodeStr, currNum)
			}
			strStack = append(strStack, decodeStr)

		} else if char == '[' || char >= 'a' && char <= 'z' {
			numTag = false
			finalNum = 0
			strStack = append(strStack, string(char))
		}
	}

	for _, s := range strStack {
		FinalDecodedString = FinalDecodedString + s
	}
	return FinalDecodedString
}
