/*
请考虑一棵二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个叶值序列 。

https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png

举个例子，如上图所示，给定一棵叶值序列为(6, 7, 4, 9, 8)的树。

如果有两棵二叉树的叶值序列是相同，那么我们就认为它们是叶相似的。

如果给定的两个根结点分别为root1 和root2的树是叶相似的，则返回true；否则返回 false 。



示例 1：
https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-1.jpg


输入：root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
输出：true
示例 2：

输入：root1 = [1], root2 = [1]
输出：true
示例 3：

输入：root1 = [1], root2 = [2]
输出：false
示例 4：

输入：root1 = [1,2], root2 = [2,2]
输出：true
示例 5：
https://assets.leetcode.com/uploads/2020/09/03/leaf-similar-2.jpg


输入：root1 = [1,2,3], root2 = [1,3,2]
输出：false


提示：

给定的两棵树可能会有1到 200个结点。
给定的两棵树上的值介于 0 到 200 之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/leaf-similar-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
  pub fn leaf_similar(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> bool {
    fn dfs(root: &Option<Rc<RefCell<TreeNode>>>,seq:&mut Vec<i32>){
      if let Some(node) = root {
        if node.borrow().left.is_none()&&node.borrow().right.is_none() {
          seq.push(node.borrow().val);
        }
        dfs(&node.borrow().left, seq);
        dfs(&node.borrow().right,seq);
      }
    }
    let mut v1 = vec![];
    let mut v2 = vec![];
    dfs(&root1,&mut v1);
    dfs(&root2,&mut v2);
    v1==v2
  }
}