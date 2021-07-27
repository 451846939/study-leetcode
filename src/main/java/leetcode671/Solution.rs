/*
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。

更正式地说，root.val = min(root.left.val, root.right.val) 总成立。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。



示例 1：
https://assets.leetcode.com/uploads/2020/10/15/smbt1.jpg

输入：root = [2,2,5,null,null,5,7]
输出：5
解释：最小的值是 2 ，第二小的值是 5 。
示例 2：

https://assets.leetcode.com/uploads/2020/10/15/smbt2.jpg
输入：root = [2,2,2]
输出：-1
解释：最小的值是 2, 但是不存在第二小的值。


提示：

树中节点数目在范围 [1, 25] 内
1 <= Node.val <= 231 - 1
对于树中每个节点 root.val == min(root.left.val, root.right.val)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree
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
    pub fn find_second_minimum_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
      if root.is_none() {
        return -1;
      }
      let mut ans=-1;
      let root_val=root.as_ref().unwrap().as_ref().borrow().val;
      fn dfs(min: &mut i32, node: &Option<Rc<RefCell<TreeNode>>>,root_val:i32){
        match node {
          None => {return;}
          Some(node) => {
            let node_val = node.as_ref().borrow().val;
            if *min != -1&& node_val >=*min {
              return;
            }
            if node_val>root_val {
              *min=node_val
            }
            dfs(min, &node.as_ref().borrow().left,root_val);
            dfs(min, &node.as_ref().borrow().right,root_val);
          }
        }
      }
      dfs(&mut ans, &root,root_val);
      ans
    }
}