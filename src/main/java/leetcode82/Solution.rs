/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中没有重复出现的数字。

返回同样按升序排列的结果链表。



示例 1：
https://assets.leetcode.com/uploads/2021/01/04/linkedlist1.jpg

输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]
示例 2：
https://assets.leetcode.com/uploads/2021/01/04/linkedlist2.jpg

输入：head = [1,1,1,2,3]
输出：[2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序排列

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
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
    let mut vir = Some(Box::new(ListNode { val: 0, next: head }));
    let mut cur = vir.as_mut();

    while cur.as_ref().unwrap().next.is_some() {
      let mut ptr = cur.as_mut().unwrap().next.as_mut();
      let val = ptr.as_mut().unwrap().val;
      let mut flag = false;
      while ptr.as_mut().unwrap().next.is_some() && ptr.as_mut().unwrap().next.as_mut().unwrap().val == val {
        ptr = ptr.unwrap().next.as_mut();
        flag = true;
      }
      if flag {
        cur.as_mut().unwrap().next = ptr.unwrap().next.take()
      } else {
        cur = cur.unwrap().next.as_mut()
      }
    }
    vir.unwrap().next
  }
}