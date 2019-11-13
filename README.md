# Go language study
## Chapter One
#### Installation
##### 1. 下载
   * 官网：<https://www.golang.org>
   * GO语言中文网:<https://studygolang.com/dl>
   * 正常安装，暂时不考虑GO_PATH的作用。
##### 2. 开发环境
   * vi,emacs,idea,eclipse,vs,sublime…… + go插件
   * IDE：Goland, LiteIDE
## Chapter two
#### Go语言的基本语法
```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```
#### 变量的定义
   * 用var关键字定义
   * 变量的类型写在变量名之后
   * 变量可以定义在函数内部，或者直接定义在包内(函数外部)，定义在函数内时，var可以省略，用`:=`代替
   * 定义在函数外部的变量的作用域为包范围而不是全局变量
   * var()集中定义变量
   * 变量可以同时定义多个，且不需要必有初始值
   * 变量定义之后必须被使用，否则报错
   * 变量类型可以省略，编译器根据赋值判断
   
   ```
   func variableValue() {
        var a,b int = 3, 4 //可同时声明多个
        var s string = "abc"
        var c = true //可省略变量类型
        d := 5 //可用 := 代替var
        fmt.Println(a, b, c, d, s)
   }
   ```
#### 内建变量类型
   * bool string
   * (u)int int8 int16 int32 int64 uintptr
      + 加u为无符号整数，不加则为有符号整数
      + 规定长度与不规定长度，不规定程度的整数根据操作系统来定义长度
      + uintptr 指针 长度根据操作系统
   * byte rune
      + rune GO语言的字符型，即char，长度为32位，byte 8位
   * float32 float64 complex64 complex128
      + complex 复数类型 1+i 
        - 复数分为实部和虚部 【泰勒级数展开】   
          ![泰勒级数展开](./images/complex.png)   
          ![泰勒级数展开](./images/complex2.png)   
          ![泰勒级数展开](./images/complex1.png)   
#### 强制类型转换
   * 类型的转换是强制的
   * 
   

