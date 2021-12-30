package main

import "fmt"

func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	test := "abc"
	fmt.Println(lengthOfNoRepeatingSubStr(test))
	test = "中国人不骗中国人"
	fmt.Println(lengthOfNoRepeatingSubStr(test))
	fmt.Println(lengthOfNoRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
