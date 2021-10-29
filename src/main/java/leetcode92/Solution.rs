/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤m≤n≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
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
    pub fn reverse_between(mut head: Option<Box<ListNode>>, left: i32, right: i32) -> Option<Box<ListNode>> {
        let mut dm = Some(Box::new(ListNode { val: 0, next: head }));
        let mut dm_mut=dm.as_mut();
        let mut i = 0;
        // let mut left_node = None;
        let mut pre = None;
        while let Some(mut left_head) = dm_mut {
            i += 1;
            while i >= left && i <= right {
                let mut cur = left_head.next.take();
                left_head.next= cur.as_mut().unwrap().next.take();
                cur.as_mut().unwrap().next= pre.take();
                if i<right {
                    pre= cur;
                }else {
                    let mut rev = cur.as_mut();
                    while let Some(rev_node) = rev {
                        if rev_node.next.is_none() {
                            rev_node.next= left_head.next.take();
                            break;
                        }
                        rev=rev_node.next.as_mut();
                    }
                    left_head.next= cur;
                    return dm.unwrap().next
                }
                i+=1;
            }
            dm_mut = left_head.next.as_mut();
        }

        dm.unwrap().next
    }

}