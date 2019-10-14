package leetcode2;

/**
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 * <p>
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 * <p>
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 * <p>
 * 示例：
 * <p>
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 * <p>
 * 来源：力扣（LeetCode）
 * 链接：https://leetcode-cn.com/problems/add-two-numbers
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
        int[] a1 = {9,9};
        ListNode oneNode = new ListNode(a1[0]);
        for (int i = 1; i < a1.length; i++) {
            int count = 1;
            for (ListNode n = oneNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a1[i]);
                }
            }
        }
        int[] a2 = {9,9};
        ListNode twoNode = new ListNode(a2[0]);
        for (int i = 1; i < a2.length; i++) {
            int count = 1;
            for (ListNode n = twoNode; count <= i; n = n.next, count++) {
                if (n.next == null) {
                    n.next = new ListNode(a2[i]);
                }
            }
        }
        addTwoNumbers(oneNode, twoNode);
    }

    public static ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode res = new ListNode(0);
        int nexttmp = 0;
        ListNode t = res;
        while (l1 != null && l2 != null) {
            int tmp = l1.val + l2.val;
            if (tmp >= 10) {
                int i = tmp + nexttmp;
                if (i >= 10) {
                    t.next = new ListNode(i % 10);
                    nexttmp = 1;
                } else {
                    t.next = new ListNode(i);
                    nexttmp = 0;
                }
//                t.next = new ListNode(tmp % 10);
//                nexttmp = 1;
            } else {
                int i = tmp + nexttmp;
                if (i >= 10) {
                    t.next = new ListNode(i % 10);
                    nexttmp = 1;
                } else {
                    t.next = new ListNode(i);
                    nexttmp = 0;
                }
            }
            l1 = l1.next;
            l2 = l2.next;
            t = t.next;
        }
        ListNode ln = l1 == null ? l2 : l1;
        if (nexttmp != 0) {
            while (ln != null) {
                int i = ln.val + nexttmp;
                if (i >= 10) {
                    t.next = new ListNode(i % 10);
                    nexttmp = 1;
                } else {
                    t.next = new ListNode(i);
                    nexttmp = 0;
                }
                ln=ln.next;
                t=t.next;
            }
        }else {
            t.next=ln;
        }
        if (nexttmp!=0){
            t.next=new ListNode(nexttmp);
        }
        return res.next;
    }
}
