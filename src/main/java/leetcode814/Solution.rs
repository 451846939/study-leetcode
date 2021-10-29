/*
给你二叉树的根结点 root ，此外树的每个结点的值要么是 0 ，要么是 1 。

返回移除了所有不包含 1 的子树的原二叉树。

节点 node 的子树为 node 本身加上所有 node 的后代。

 

示例 1：
https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_2.png

输入：root = [1,null,0,0,1]
输出：[1,null,0,null,1]
解释：
只有红色节点满足条件“所有不包含 1 的子树”。 右图为返回的答案。
示例 2：
https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/06/1028_1.png

输入：root = [1,0,1,0,0,0,1]
输出：[1,null,1,null,1]
示例 3：

https://s3-lc-upload.s3.amazonaws.com/uploads/2018/04/05/1028.png
输入：root = [1,1,0,1,1,0,1,0]
输出：[1,1,0,1,1,null,1]
 

提示：

树中节点的数目在范围 [1, 200] 内
Node.val 为 0 或 1


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-pruning
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
      right: None,
    }
  }
}

use std::rc::Rc;
use std::cell::RefCell;

pub fn prune_tree(mut root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {

  fn contains_one(node: Option<&mut Rc<RefCell<TreeNode>>>) ->bool{
    return if let Some(n) = node {
      let mut node = n.borrow_mut();
      let left = contains_one(node.left.as_mut());
      let right = contains_one(node.right.as_mut());
      if !left {
        node.left = None;
      }
      if !right {
        node.right = None;
      }
      node.val == 1 || left || right
    } else {
      false
    }
  }
  if !contains_one(root.as_mut()) {
    return None;
  }
  root
}