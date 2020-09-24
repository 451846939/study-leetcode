package main

/*
给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

结点左子树中所含结点的值小于等于当前结点的值
结点右子树中所含结点的值大于等于当前结点的值
左子树和右子树都是二叉搜索树
例如：
给定 BST [1,null,2,2],

   1
    \
     2
    /
   2
返回[2].

提示：如果众数超过1个，不需考虑输出顺序

进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-mode-in-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func main() {

}

/**
 * Definition for a binary tree node. */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) (res []int) {
	var base, count, maxCount int
	update := func(val int) {
		if val == base {
			count++
		} else {
			base, count = val, 1
		}
		if count == maxCount {
			res = append(res, base)
		} else if count > maxCount {
			maxCount = count
			res = []int{base}
		}
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)

	//以下为 Morris  中序遍历
	//cur := root
	//for cur != nil {
	//	if cur.Left == nil {
	//		update(cur.Val)
	//		cur = cur.Right
	//		continue
	//	}
	//	pre := cur.Left
	//	for pre.Right != nil && pre.Right != cur {
	//		pre = pre.Right
	//	}
	//	if pre.Right == nil {
	//		pre.Right=cur
	//		cur=cur.Left
	//	}else {
	//		pre.Right=nil
	//		update(cur.Val)
	//		cur=cur.Right
	//	}
	//}
	return
}
