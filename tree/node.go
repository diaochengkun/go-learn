package tree

import "fmt"

/**
演示go的封装，注意go语言没有  继承  跟 多态
 */

type  Treenode struct {
	Value int
	Left,Right *Treenode
}

//类似于java 中类的 方法定义
func  (node *Treenode) print()  {
	fmt.Println(node.Value)
}

//二叉树的遍历
func (node *Treenode) Traverse()  {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.print()
	node.Right.Traverse()
}

func (node *Treenode) SetValue(value int)  {
	node.Value = value
}


