package main

import "fmt"

func lengthOfNoRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		//fmt.Printf("i=%d,ch=%c,lastOccurred[ch]=%d,start=%d,maxLength=%d\n", i, ch, lastOccurred[ch], start, maxLength)
		lastI, ok := lastOccurred[ch]
		//fmt.Printf("lastI=%d,ok=%t\n", lastI, ok)
		if ok && lastI >= start {
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
}
