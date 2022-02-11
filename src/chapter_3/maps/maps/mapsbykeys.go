package main

import "fmt"

type queryKey struct {
	Name string
	Age  int
}
type Profile struct {
	Name    string
	Age     int
	Address string
}

var mapper = make(map[queryKey]*Profile)

func buildIndex(list []*Profile) {
	for _, profile := range list {
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}
		fmt.Println(key)
		mapper[key] = profile
	}
}

func queryData(name string, age int) {
	key := queryKey{Name: name, Age: age}
	result, ok := mapper[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("没有找到对应数据")
	}
}

func findData(list []*Profile, name string, age int) {
	for _, data := range list {
		if data.Name == name && data.Age == age {
			fmt.Println(data)
			return
		}
	}
	fmt.Println("没有找到对应的数据")
}

func main() {
	list := []*Profile{
		{Name: "张三", Age: 23, Address: "cq"},
		{Name: "李四", Age: 25},
		{Name: "王五"},
	}
	//传统查询
	findData(list, "张三", 23)
	//利用map的多键索引查询（组合键查询）
	buildIndex(list)
	queryData("李四", 25)
}
