package main

import (
	"strconv"
	"strings"
)

/*

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode(https://support.leetcode-cn.com/hc/kb/category/1018267/) 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {

}

/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	l []string
}

func Constructor() Codec {
	return Codec{make([]string, 0)}
}
func dfsSerialize(root *TreeNode, s string) string {
	if root == nil {
		s += "null,"
	} else {
		s += strconv.Itoa(root.Val) + ","
		s = dfsSerialize(root.Left, s)
		s = dfsSerialize(root.Right, s)
	}
	return s
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return dfsSerialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for i := range l {
		if l[i] != "" {
			this.l = append(this.l, l[i])
		}
	}
	root := this.dfsDeserialize()

	return root
}

func (this *Codec) dfsDeserialize() *TreeNode {
	if this.l[0] == "null" {
		this.l = this.l[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.l[0])
	root := &TreeNode{Val: val}
	this.l = this.l[1:]
	root.Left = this.dfsDeserialize()
	root.Right = this.dfsDeserialize()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
