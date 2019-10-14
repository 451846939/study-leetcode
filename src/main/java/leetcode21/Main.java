package leetcode21;

/**
 * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
 * <p>
 * 示例：
 * <p>
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
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
        int[] a1 = {1, 2, 3, 4, 5};
        ListNode oneNode = new ListNode(a1[0]);
        for (int i = 1; i < a1.length; i++) {
            int count = 1;
            for (ListNode n = oneNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a1[i]);
                }
            }
        }
        int[] a2 = {1, 2, 3, 4, 5};
        ListNode twoNode = new ListNode(a2[0]);
        for (int i = 1; i < a2.length; i++) {
            int count = 1;
            for (ListNode n = twoNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a2[i]);
                }
            }
        }
        ListNode listNode = mergeTwoLists(oneNode, twoNode);
        System.out.println(listNode);
    }

    public static ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode res = new ListNode(0);
        ListNode restmp = res;
        while (l1 != null && l2 != null) {
            if (l1.val > l2.val) {
                restmp.next = l2;
                l2 = l2.next;
            } else {
                restmp.next = l1;
                l1 = l1.next;
            }
            restmp = restmp.next;
        }
        if (l1==null){
            restmp.next=l2;
        }
        if (l2==null){
            restmp.next=l1;
        }
        return res.next;
    }
    public static ListNode mergeTwoLists2(ListNode l1, ListNode l2) {
        // maintain an unchanging reference to node ahead of the return node.
        ListNode prehead = new ListNode(-1);

        ListNode prev = prehead;
        while (l1 != null && l2 != null) {
            if (l1.val <= l2.val) {
                prev.next = l1;
                l1 = l1.next;
            } else {
                prev.next = l2;
                l2 = l2.next;
            }
            prev = prev.next;
        }

        // exactly one of l1 and l2 can be non-null at this point, so connect
        // the non-null list to the end of the merged list.
        prev.next = l1 == null ? l2 : l1;

        return prehead.next;
    }
}
