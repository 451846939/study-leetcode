/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。



示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}
impl Solution {
    pub fn delete_duplicates(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dum =Some(Box::new(ListNode{val:0,next:head}));
        let mut cur=dum.as_mut();
        while cur.as_ref().unwrap().next.is_some() {
            let mut ptr =cur.as_mut().unwrap().next.as_mut();
            if ptr.as_mut().unwrap().next.is_some()&&ptr.as_mut().unwrap().val==ptr.as_mut().unwrap().next.as_mut().unwrap().val {
                cur.as_mut().unwrap().next=ptr.unwrap().next.take()
            }else {
                cur=cur.unwrap().next.as_mut()
            }
        }
        dum.unwrap().next
    }
}