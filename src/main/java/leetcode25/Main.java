package leetcode25;

/**
 * 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
 * <p>
 * k 是一个正整数，它的值小于或等于链表的长度。
 * <p>
 * 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
 * <p>
 * 示例 :
 * <p>
 * 给定这个链表：1->2->3->4->5
 * <p>
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 * <p>
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 * <p>
 * 说明 :
 * <p>
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
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
        int[] a1 = {1, 2,3,4};
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
        reverseKGroup(oneNode, 3);
    }

    public static ListNode reverseKGroup(ListNode head, int k) {
        if (head==null){
            return head;
        }
        ListNode res = new ListNode(0);
        res.next = head;
        ListNode prev = res;
        ListNode[] listNodes = new ListNode[k];
        ListNode lastNode;
        while (true){
            lastNode=prev;
            for (int i = 0; i < k; i++) {
                if (prev!=null){
                    listNodes[i] = prev.next;
                    prev = prev.next;
                }
            }
            if (listNodes[k-1]==null){
                break;
            }
            for (int i=0; i<k;i++){
                if (i==0){
                    listNodes[i].next=prev.next;
                    continue;
                }
                listNodes[i].next=listNodes[i-1];
                if (i==k-1){
                    lastNode.next=listNodes[i];
                    break;
                }
            }
            prev=listNodes[0];
            for (int i = 0; i < k; i++) {
                listNodes[i] = null;
            }
        }
        return res.next;
    }
}
