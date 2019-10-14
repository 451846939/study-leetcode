package leetcode92;

/**
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 * <p>
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 * <p>
 * 示例:
 * <p>
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
 * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
public class Main {
    public static class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
        }
    }

    public static void main(String[] args) {
        int[] a = {1, 2, 3, 4, 5};
        ListNode oneNode = new ListNode(a[0]);
        for (int i = 1; i < a.length; i++) {
            int count = 1;
            for (ListNode n = oneNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a[i]);
                }
            }
        }
        reverseBetween(oneNode, 1, 1);
    }

    public static ListNode reverseBetween(ListNode head, int m, int n) {
        if (m==n){
            return head;
        }
        ListNode prev=null;
        ListNode reshead=head;
        ListNode curr=head;
        ListNode start=head;
        ListNode startnext=null;
        ListNode end=null;
        ListNode endlast=null;
        for (int i = 1; curr!=null; i++) {
            if (i<m-1){
                curr=curr.next;
                continue;
            }
            if (i==m-1){
                start=curr;
                curr=curr.next;
                continue;
            }
            if (i>=m&&i<=n){
                if (i==m){
//                     prey  curr  next
//                     0     0    0
                    endlast=curr;
                    ListNode tmp = curr.next;
                    curr.next=prev;
                    prev=curr;
                    curr=tmp;
                }else if (i==n){
                    startnext=curr;
                    end=curr.next;
                    endlast.next=end;
                    ListNode tmp = curr.next;
                    curr.next=prev;
                    prev=curr;
                    curr=tmp;
                    if (m==1){
                        start=startnext;
                    }else {
                        start.next=startnext;

                    }
                }else {
                    ListNode tmp = curr.next;
                    curr.next=prev;
                    prev=curr;
                    curr=tmp;
                }
            }else {
                prev=curr;
                curr=curr.next;
            }
        }
        if (m==1){
            reshead=start;
        }
        return reshead;
    }
}
