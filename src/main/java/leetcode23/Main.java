package leetcode23;

/**
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 *
 * 示例:
 *
 * 输入:
 * [
 *   1->4->5,
 *   1->3->4,
 *   2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
 *
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
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
        ListNode listNode = mergeKLists(new ListNode[]{oneNode, twoNode});
        System.out.println(listNode);

    }
    public static ListNode mergeKLists(ListNode[] lists) {
        if (lists.length==1){
            return lists[0];
        }
        ListNode res=null;
        for (int i = 0; i+1 < lists.length; i++) {
            if (res==null){
                res=lists[i];
            }
            res = mergeTwoLists(res, lists[i + 1]);
        }
        return res;
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
}
