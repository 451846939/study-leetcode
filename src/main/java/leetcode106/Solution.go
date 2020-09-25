package main

/*

根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	idxMap := make(map[int]int)
	for i, v := range inorder {
		idxMap[v] = i
	}
	var build func(inorderL, inorderR int) *TreeNode
	build = func(inorderLIdx, inorderRIdx int) *TreeNode {
		if inorderLIdx > inorderRIdx {
			return nil
		}
		// 后序遍历的末尾元素即为当前子树的根节点
		val := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := &TreeNode{Val: val}

		// 根据 val 在中序遍历的位置，将中序遍历划分成左右两颗子树
		// 由于我们每次都从后序遍历的末尾取元素，所以要先遍历右子树再遍历左子树
		inorderRootIndex := idxMap[val]
		root.Right = build(inorderRootIndex+1, inorderRIdx)
		root.Left = build(inorderLIdx, inorderRootIndex-1)
		return root
	}
	return build(0, len(inorder)-1)
}

//https://pic.leetcode-cn.com/ac050d257073f47285353d7ad412fb832326237ea85948a8b69d338171d67543-%E6%A0%91%E7%9A%84%E8%BF%98%E5%8E%9F.png
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	i := 0
	for postorder[len(postorder)-1] != inorder[i] {
		i++
	}
	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}
