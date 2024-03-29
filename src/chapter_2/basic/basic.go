package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//var aa = 3
//var ss = "KKK"
//var (
//	variblesx  int
//	slicex     []int
//	interfacex interface{}
//)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	a, b = b, a
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, _, s = 3, 4, true, "def"
	fmt.Println(a, b, true, s)
}

func variableShorter() {
	//第一次使用变量可以使用冒号定义，但在函数外部定义只能使用var
	a, b, _, s := 3, 4, true, "def"
	fmt.Println(a, b, true, s)
}

//复数用例
func euler() {
	//fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

//强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举类型
func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	//b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	euler()
	triangle()

	consts()
	enums()
}
