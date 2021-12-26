package main

import "fmt"

func lengthOfNoRepeatingSubStr(s string) bool {
	lastOccurred := make(map[byte]int)
	//start := 0
	//maxLength := 0
	_, ok := lastOccurred[0]
	return ok
	//for i, ch := range []byte(s) {
	//	if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
	//		start = lastI + 1
	//	}
	//	if i-start+1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i
	//}
	//return maxLength
}

func main() {
	test := "1111"
	fmt.Println(lengthOfNoRepeatingSubStr(test))
}
