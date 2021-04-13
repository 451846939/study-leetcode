/*
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同



示例 1：


输入：root = [4,2,6,1,3]
输出：1
示例 2：


输入：root = [1,0,48,null,null,12,49]
输出：1


提示：

树中节点数目在范围 [2, 100] 内
0 <= Node.val <= 105


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes
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
    pub fn min_diff_in_bst(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
      let mut ans = std::i32::MAX;
      let mut pre = -1;
      fn dfs(node: Option<Rc<RefCell<TreeNode>>>, pre:&mut i32, ans: &mut i32) {
        if let Some(node) = node {
          let node = node.borrow();
          dfs(node.left.clone(),pre,ans);
          if *pre!=-1&&node.val-*pre<*ans {
            *ans=node.val-*pre
          }
          *pre=node.val;
          dfs(node.right.clone(),pre,ans)
        }
      }
      ;
      dfs(root, &mut pre, &mut ans);
      ans
    }
}