package main

import "fmt"

/**
此类演示数组的使用
数组属于数值类型  方法传递时会发生值拷贝
*/

func printArray1(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func printArray2(arr [5]int) {

	//rang关键字可以返回俩值，下标 值
	for i, v := range arr {
		fmt.Println(i, v)
	}

	//如果不想要下标  可以使用 _ 表示
	for _, v := range arr {
		fmt.Println(v)
	}
}
func changeArray(arr *[5]int) {
	arr[0] = 1000
}

func main() {

	//数组的定义
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	arr2 := [...]int{4, 5, 6, 7, 8}

	arr3 := [5]int{1, 2, 3, 4, 5}

	printArray1(arr1)

	printArray1(arr2)

	printArray2(arr3)

	changeArray(&arr1)
	fmt.Println(arr1)
}
