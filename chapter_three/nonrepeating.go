package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		fmt.Println(i)
		fmt.Println(ch)
		lastI, ok := lastOccurred[ch]
		fmt.Println(ok)
		fmt.Println(lastI)
		if ok && lastI >= start {
			start = lastI + 1
		}
		fmt.Println(start)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		fmt.Println(maxLength)
		fmt.Println(i)
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
}
