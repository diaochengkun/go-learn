package main

import "fmt"

func main() {

	m := map[string]string {
		"name": "zhangsan",
		"courseName": "go",
		"age": "18",
	}

	m1 := make(map[string]string)

	var m2 map[string]string

	fmt.Println(m,m1,m2)

	//map结构遍历
	for k,v := range m {
		fmt.Println(k,v)
	}
	//map 查询 第一个为value值，第二个为是否查询到
	courseName,ok := m["courseName"]
	fmt.Println(courseName,ok)

	//key不存在时 会返回空值,此时可以使用ok判断是否包含key
	if courseName1 ,ok:= m["courseName1"] ; ok {
		fmt.Println(courseName1,ok)
	} else {
		fmt.Println("此map中不存在该key")
	}

	//map删除某个key
	delete(m,"age")
	fmt.Println(m)



}
