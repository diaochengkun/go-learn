package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//展示循环用法
//go 没有 while 只有for
// for循环不需要带括号

//for 可以省略起始条件
func convertToBin(a int) string {
	result := ""
	if a == 0 {
		return "0"
	}
	for ; a > 0; a = a / 2 {
		mod := a % 2
		result = strconv.Itoa(mod) + result
	}
	return result
}

// for循环省略起始条件  结束条件     相当于while循环
func printFile(filename string) {
	file, error := os.Open(filename)
	if error != nil {
		panic(error)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//for循环可以省略起始条件  结束条件  递增表达式
// 代表死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}
func main() {
	println(
		convertToBin(12),
		convertToBin(13),
		convertToBin(123213123),
		convertToBin(0),
	)
	printFile("abc.txt")

	forever()
}
