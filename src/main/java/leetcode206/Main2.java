package leetcode206;

public class Main2 {
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
        ListNode prev=null;
        ListNode curr=head;
        while (curr!=null){
//            获取当前节点的下一个节点
            ListNode temp=curr.next;
//            把上一个节点赋值给当前循环节点的下一个节点
            curr.next=prev;
//            把当前节点赋值给上一个节点值
            prev=curr;
//            把当前节点的下一个节点赋值给当前循环节点
            curr=temp;
        }
        return prev;
    }
}
