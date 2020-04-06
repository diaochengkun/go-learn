package main

import (
	"fmt"
	"math"
)

var aaa = "aa"
var bbb = "bbb"

//可以定义全局变量，作用域为包内
//全局变量必须以关键字开头，因此定义变量不得 用冒号:了
//为了省略 多个 var  可以使用括号代替
var (
	ccc = "ccc"
	ddd = "ddd"
)

//int 默认0 string默认空串
func variableZero() {
	var a int
	var s string
	fmt.Println(a, s)
}

func variableInitival() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d, e = 1, 2, 3, "abc", true

	fmt.Println(a, b, c, d, e)
}

//变量定义可以不使用var  :=代替
// 之后使用不带冒号
func variableShort() {
	a, b, c, d, e := 1, 2, 3, "abc", true
	b = 5
	fmt.Println(a, b, c, d, e)
}

//演示强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//定义变量 使用关键字 const
func constant() {
	const filename = "abc.txt"
	const a, b = 1, 2
	fmt.Println(filename, a, b)
}

//定义枚举
func enums() {
	const (
		cpp        = 1
		java       = 2
		goo        = 3
		javascript = 4
	)
	fmt.Println(cpp, java, goo, javascript)

	//iota表示0 后续常量自增
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, tb, pb)
}
func main() {
	variableZero()
	variableInitival()
	variableTypeDeduction()
	variableShort()
	triangle()
	constant()
	enums()
}
