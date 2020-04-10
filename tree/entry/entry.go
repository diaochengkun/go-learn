package main

import "golearn/tree"

/**
go中包 和 封装

1.go也有package概念  package可以不用跟目录保持一致
2.同一目录下 package必须相同

首字母大写 表示 public
首字母小写 表示private
 */
func main() {
	//结构体初始化几种方式
	root := tree.Treenode{Value: 2}

	root.Left = &tree.Treenode{}
	root.Left.Value = 5
	root.Right = new(tree.Treenode)
	root.Right.SetValue(3)

	root.Traverse()
}
