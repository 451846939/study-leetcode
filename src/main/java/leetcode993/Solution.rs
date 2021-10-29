/*
在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但 父节点不同 ，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root ，以及树中两个不同节点的值 x 和 y 。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true 。否则，返回 false。



示例 1：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/q1248-01.png

输入：root = [1,2,3,4], x = 4, y = 3
输出：false
示例 2：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/q1248-02.png

输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
输出：true
示例 3：
https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2019/02/16/q1248-03.png


输入：root = [1,2,3,null,4], x = 2, y = 3
输出：false


提示：

二叉树的节点数介于2 到100之间。
每个节点的值都是唯一的、范围为1 到100的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cousins-in-binary-tree
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
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
      type Tree = Option<Rc<RefCell<TreeNode>>>;
      fn find(root: Tree, x: i32, p: Tree, l: i32) -> Option<(Tree, i32)> {
        let l1 = l + 1;
        if let Some(node) = &root {
          let node = node.borrow();
          if node.val == x {
            return Some((p, l1));
          } else {
            if let Some(r) = find(node.left.clone(), x, root.clone(), l1) {
              return Some(r);
            }
            if let Some(r) = find(node.right.clone(), x, root.clone(), l1) {
              return Some(r);
            }
          }
        }
        None
      }

      let (f1, l1) = find(root.clone(), x, None, 0).unwrap();
      let (f2, l2) = find(root.clone(), y, None, 0).unwrap();
      f1 != f2 && l1 == l2
    }
}