package chapter_two

import "fmt"

var aa = 3
var ss = "KKK"

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//第一次使用变量可以使用冒号定义，但在函数外部定义只能使用var
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}
func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
