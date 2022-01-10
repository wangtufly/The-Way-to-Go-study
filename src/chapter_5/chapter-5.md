# Go language study

## chapter 5

### 1. 接口

#### duck typing

* 描述事物的外部行为而非内部结构 长得像鸭子就是鸭子
* 严格说GO属于结构化类型系统，类似duck typing

#### 接口的定义

* 接口由使用者定义

> 接口变量里面有什么 接口变量[实现者类型，实现者指针]
> * 接口变量自带指针
> * 接口变量同样采用值传递，几乎不需要使用接口的指针
> * 指针接收者实现只能以指针方式使用；值接收者都可

#### 查看接口变量

* 表示任何类型：interface{}
* Type Assertion
* Type Switch

#### 接口组合
```
type ReadWriter interface
{
  Reader
  Writer
}
```


