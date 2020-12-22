/*
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
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
use std::collections::{VecDeque};
impl Solution {
  pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
    let mut res: Vec<Vec<i32>> = vec![];
    if root.is_none() {return res }
    let mut q=VecDeque::new();
    q.push_back(root);
    let mut backward = false;
    while q.len()>0 {
      let mut temp = vec![];
      for _ in 0..q.len() {
        if let Some(node)=q.pop_front().unwrap() {
          temp.push(node.borrow().val);
          let left = node.borrow_mut().left.take();
          let right = node.borrow_mut().right.take();
          if left.is_some() { q.push_back(left); }
          if right.is_some() { q.push_back(right); }
        }
      }
      if backward { temp.reverse(); }
      backward = !backward;
      res.push(temp);
    }
    return res;
  }


}