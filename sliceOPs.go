package main

import "fmt"

/**
slice的操作
 */

func  printSlice(arr []int)  {
	fmt.Printf("%v,lens=%d,cap=%d\n",arr,len(arr),cap(arr))
}
func main() {
	var arr []int

	//添加元素
	for i := 0;i< 100;i++{
		printSlice(arr)
		arr = append(arr,2* i +1)
	}

	s1 := []int{0,1,2,3}

	printSlice(s1)

	//长度是16
	s2 := make([]int,16)

	printSlice(s2)

	//长度是 10  容量是 32
	s3 := make([]int,10,32)
	printSlice(s3)


	//copy

	copy(s2,s1)
	printSlice(s2)

	//delete index fro slice
	//
	s2 = append(s2[:3],s2[4:]...)
	printSlice(s2)

	fmt.Println("poping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println("poping from tail")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("front=%d,taild=%d",front,tail)
	printSlice(s2)
}
