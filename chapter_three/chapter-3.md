# Go language study

## chapter three

### 1. 数组，切片和容器

#### 1.1 数组

* `arr [len]int`/`arr:=[len]int{}`/`arr:=[...]int{}`
* `arr [num][len]int`
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
s[0]=2
```

* s的值为`[2 3 4 5]`
* slice 本事没有数据，是对底层array的一个view
* arr的值变为`[0 1 10 3 4 5 6 7]`
* slice的扩展

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
* map的遍历
    + 使用range 遍历key 或者遍历key value对
    + 不保证遍历顺序，如需顺序，需要手动对key排序
    + 使用len获取元素 个数
* map的key
    + map使用哈希表，必须可以比较相等
    + 除了slice，map，function的内建类型都可以作为key
    + struct类型不包含上述字段，也可以做key
* 例：寻找最长不含有重复字符的子串
