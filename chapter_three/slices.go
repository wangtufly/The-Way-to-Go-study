package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[:6]", arr[:6])
	fmt.Println("arr[2:]", arr[2:])
	fmt.Println("arr[:]", arr[:])
	s1 := arr[3:]
	fmt.Println("s1", s1)
	s2 := arr[:]
	fmt.Println("s2", s2)
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println("s2的切片", s2)
	s2 = s2[2:]
	fmt.Println("s2切片的切片", s2)

	fmt.Println("extending slice")
	arr[0], arr[3] = 0, 3
	s1 = arr[2:6]
	s2 = s1[3:5] //[s1[3],s1[4]]
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	//fmt.Println(s1[3:8])//slice bounds out of range [:8] with capacity 6

	s3 := append(s2, 11)
	s4 := append(s3, 12)
	s5 := append(s4, 13)
	fmt.Println("s3,s4,s5", s3, s4, s5)
	fmt.Println("arr=", arr)
}
