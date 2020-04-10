package main

import "fmt"

/**
演示切片的概念及使用
 */

func updateSlice(arr []int)  {
	arr[0] = 100
}
func main() {

	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := arr[:6]
	s3 := arr[:]
	fmt.Printf("s1=%v,s2=%v,s3=%v\n",s1,s2,s3)
	fmt.Println("update slice.....")
	updateSlice(s1[:])
	fmt.Printf("arr=%v,s1=%v\n",arr,s1)

	/*
	切片是原始数组的view  记录三部分
	1: ptr 原始数组的引用
	2.len 切片数组的长度
	3.cap 原始数组的容量
	 */
	arr = [...]int{0,1,2,3,4,5,6,7}
	s1 = arr[2:6]
	s2 = s1[3:5]
	fmt.Printf("s1=%v,s2=%v\n",s1,s2)


	//添加元素时如果超过cap，系统会重新分配更大的底层数组
	s4 := append(s2,10)
	s5 := append(s4,11)
	s6 := append(s5,12)
	fmt.Printf("s3=%v,s4=%v,s5=%v\n",s4,s5,s6)
	fmt.Println("arr = ",arr)


	
}
