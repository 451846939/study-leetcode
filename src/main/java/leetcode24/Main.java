package leetcode24;

/**
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *  
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
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
//        int[] a1 = {1,2,3,4,6};
        int[] a1 = {0};
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
        swapPairs(oneNode);
    }
    public static ListNode swapPairs(ListNode head) {
        ListNode res = new ListNode(0);
        res.next=head;
        ListNode prev = res;
        ListNode n1;
        ListNode n2;
        ListNode next;
        while (prev.next!=null&&prev.next.next!=null){
                n1=prev.next;
                n2=prev.next.next;
                next=n2.next;
                n1.next=next;
                n2.next=n1;
                prev.next=n2;
                prev=n1;
        }
        return res.next;
    }
}
