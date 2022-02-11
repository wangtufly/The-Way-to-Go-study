# Go language study

## chapter three

### 1. 数组，切片和容器

#### 1.1 数组

* `arr [len]int`/`arr:=[len]int{}`/`arr:=[...]int{}`
* `arr [num][len]int` 矩阵数组
* 数量写在类型前面
* 数组的遍历--range关键字:
  ```
  for _,v := range numbers{
  }
  ```
    + 可通过`_`省略变量
    + 不仅range，任何地方都可以用`_`省略变量
    + 如果只要i,可以写成`for i:=range numbers`
  > 为什么要用range
  >    - 意义明确，美观
  >    - c++: 没有类似能力
  >    - java/python: 只能foreach value,不能同时获取i，v

* 数组是值类型
    + `[10]int`和`[20]int`是不同类型
    + 调用`func f(arr [10]int)`会拷贝数组
    + 在go语言中一般不直接使用数组

#### 1.2 切片slice

```
arr:=[...]int{0,1,2,3,4,5,6,7}
s:=arr[2:6]
s[0]=12
```

* s的值为`[2 3 4 5]`
* slice 本事没有数据，是对底层array的一个view
* 由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。[前闭后开)
* arr的值变为`[0 1 12 3 4 5 6 7]`
* slice的扩展 reslice

```
arr:=[...]int{0,1,2,3,4,5,6,7}
s1:=arr[2:6]
s2:=s1[3:5]
```

结果是`s1=[2,3,4,5]，s2=[5,6]`

* slice的实现
    + slice 的结构中含有 ptr，len，cap三个参数
    + slice可以向后扩展，不能向前扩展。
    + `s[i]`不可超越len(s)，向后扩展不可以超越底层数组cap(s)
* 向slice添加元素
    + 添加元素时如果超越cap,系统会重新分配更大的底层数组
    + 由于值传递的关系，必须接受append的返回值
    + s=append(s,val)
* 删，建
    + slice copy 类似于php 数组相加
    + 直接用`[:]`截取
    + `make([]int,len,cap)`

#### 1.3 Map

```
m:= map[string]string{
    "name":"ccmouse",
    "course":"golang",
    "site":"imooc",
    "quality":"notbad",
}
```

* map定义:`map[K]V`,`map[K1]map[K2]V`
    + 创建：`make(map[string]int)`
    + 获取元素:`m[key]`
    + key不存在时，获得Value类型的初始值
    + 用 `value,ok:=m[key]`来判断是否存在key
    + delete删除一个key
    + map不用手动清空，会自动回收，所以需要一般直接写新map
* map的遍历
    + 使用range 遍历key 或者遍历key value对
    + 不保证遍历顺序，如需顺序，需要手动对key排序
    + 使用len获取元素 个数
* map的key
    + map使用哈希表，必须可以比较相等
    + 除了slice，map，function的内建类型都可以作为key
    + struct类型不包含上述字段，也可以做key
* 例：寻找最长不含有重复字符的子串
* 基于map的多键值索引，可以提高查询效率
* map在并发的情况下，只读线程安全，同时读写线程不安全
    + 并发读写，map内部会对此进行检查并提前发现
    + 需要并发读写时，一般做法是加锁，但效率不高，go在>=1.9中提供了高效并发安全的sync.Map,但不是以语言原生形态提供，而是在sync包下的特殊结构
* sync.Map
    + 无须初始化，直接声明即可
    + sync.Map不能使用map的方式进行取值和设置操作，而是使用sync.Map的方法进行调用，Store表示存储，Load表示获取，Delete表示删除
    + 使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，range参数中回调函数的返回值在需要继续迭代遍历时，返回true，终止迭代遍历时，返回false
    + sync.Map没有提供获取map数量的方法，且有一定性能损失，在非并发的情况下，使用map更高效

### 1.4 rune

* rune相当于go的char
    + 使用range遍历pos,rune对（i 英文连续 中文+3）
    + 使用utf.RuneCountInString获得字符数量
    + 使用len获得字节长度
    + 使用[]byte获得字节
* 其他字符串操作
    + Fields,Split,join
    + Contains,index
    + ToLower,ToUpper
    + Trim,TrimRight,TrimLeft

### 1.5 list
* 单链表，双链表
* 列表使用container/list包来实现，内部实现原理是双链表
* list 的初始化分别是New()函数和var关键字声明，两种方法效果一致
* 列表与切片和map不同的是，列表并没有具体的元素类型限制
