package leetcode206;

import java.util.ArrayList;

public class Main {
  public static class ListNode {
      int val;
      ListNode next;
      ListNode(int x) { val = x; }
  }

    public static void main(String[] args) {
        int[] a={1,2,3,4,5};
        ListNode oneNode = new ListNode(a[0]);
        for (int i = 1; i < a.length; i++) {
            int count=1;
            for (ListNode n=oneNode;count<=i;n=n.next,count++){
                if (n.next==null){
                    n.next=new ListNode(a[i]);
                }
            }
        }
        reverseList(oneNode);
    }
    public static ListNode reverseList(ListNode head) {
        if (head==null||head.next==null){
            return head;
        }
        ArrayList<Integer> list=new ArrayList();
        for(ListNode node=head;node.next!=null;node=node.next){
            list.add(node.val);
            if (node.next.next==null){
                list.add(node.next.val);
            }
        }
        ListNode firstNode=new ListNode(list.get(list.size()-1));
        for (int i = list.size()-2; i >=0 ; i--) {
            int count=list.size()-2;
            for (ListNode n=firstNode;count>=i;n=n.next,count--){
                if (n.next==null){
                    n.next=new ListNode(list.get(i));
                }
            }
        }
        return firstNode;
    }
}
