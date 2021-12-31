# Go language study

## chapter four

### 1. 面向对象

* go语言仅支持封装，不支持继承和多态
* go语言没有class，只有struct
* 不论地址还是结构本身，一律使用.来访问成员

```
    root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
```

#### 结构的创建

* 使用自定义工厂函数

```
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
root.left.right = createNode(2)
```

* 注意返回了局部变量的地址

> 问 结构创建在堆上还是栈上？ 不需要知道

#### 为结构定义方法

* 显式定义和命名方法接收者

```
func (node treeNode) print() {
	fmt.Print(node.value, "\n")
}
```

* 只有使用指针才能改变结构内容

```
func (node *treeNode) setValue(value int) {
	node.value = value
}
```

* nil指针也可以调用方法

#### 值接收者 vs 指针接收者

* 要改变内容必须使用指针接收者
* 结构过大也考虑使用指针接收者
* 一致性:如有指针最好用指针
* 值接收者是GO特有
* 值/指针接收者均可接收值/指针

#### 封装

* 名字一般使用CamelCase 大驼峰
* 首字母大写 public 针对包
* 首字母小写 private 针对包

#### 包

* 每个目录一个包
* main包包含可执行入口
* 为结构定义的方法必须放在同一个包内
* 可以是不同的文件
> 如何过充系统类型或者别人的类型
> 定义别名
> 使用组合

#### GOPATH 下目录结构
* go build 来编译
* go install 产生pkg文件和可执行文件
* go run 直接编译运行

