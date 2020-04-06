package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

//if 的用法
func iftest() {
	filename := "abc.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

//if的简写用法
func iftest1() {
	filename := "abc.txt"
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

//switch 用法  case后面不用带break
func switchtest(a, b int, ope string) int {
	var result int
	switch ope {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator " + ope)
	}
	return result
}

//switch 用法 2  switch 后面可以不用 加 表达式
func switchtest1(grade int) string {
	var result string
	switch {
	case grade < 0 || grade > 100:
		panic(" illegal grade " + strconv.Itoa(grade))
	case grade <= 60:
		result = "F"
	case grade <= 80:
		result = "C"
	case grade <= 90:
		result = "B"
	case grade <= 100:
		result = "A"
	}
	return result
}
func main() {
	iftest()
	iftest1()
	fmt.Println(switchtest(1, 12, "+"))
	fmt.Println(switchtest1(12))
	fmt.Println(
		switchtest1(0),
		switchtest1(60),
		switchtest1(80),
		switchtest1(90),
		switchtest1(100),
		switchtest1(123),
	)
}
