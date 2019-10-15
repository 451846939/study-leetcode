package leetcode19;

/**
 *给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
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
        int[] a1 = {1, 2,3,4,5};
//        int[] a1 = {0};
        ListNode oneNode = new ListNode(a1[0]);
        for (int i = 1; i < a1.length; i++) {
            int count = 1;
            for (ListNode n = oneNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a1[i]);
                }
            }
        }
//        int[] a2 = {9,9};
//        ListNode twoNode = new ListNode(a2[0]);
//        for (int i = 1; i < a2.length; i++) {
//            int count = 1;
//            for (ListNode n = twoNode; count <= i; n = n.next, count++) {
//                if (n.next == null) {
//                    n.next = new ListNode(a2[i]);
//                }
//            }
//        }
        removeNthFromEnd(oneNode, 2);
    }

    public static ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode res = new ListNode(0);
        res.next=head;
        ListNode p = res;
        ListNode q;
        ListNode h=res.next;
        int i=1;
        boolean f=true;
        while (h!=null){
            if (i<=n){
                q=h;
                i++;
            }else {
                q=h.next;
                p=p.next;
            }
            h=h.next;
            if (q==null){
                f=false;
                p.next=p.next.next;
                break;
            }
        }
        if (f){
            p.next=p.next.next;
        }
        return res.next;
    }

    public static ListNode removeNthFromEnd2(ListNode head, int n) {
        ListNode res = new ListNode(0);
        res.next=head;
        ListNode p = res;
        ListNode q=res;
        for (int i = 0; i <=n; i++) {
            q=q.next;
        }
        while (q!=null){
            q=q.next;
            p=p.next;
        }
        p.next=p.next.next;
        return res.next;
    }
}
