package structures;

import java.util.List;

public class ListNode {

  public int val;
  public ListNode next;

  public ListNode(int v) {
    val = v;
  }

  // 创建链表
  public static ListNode createListNode(List<Integer> values) {
    ListNode node = new ListNode(-1);
    var dummy = node;
    for (int i = 0; i < values.size(); i++) {
      node.next = new ListNode(values.get(i));
      node = node.next;
    }
    return dummy.next;
  }

  // 打印链表
  public static void print(ListNode listNode) {
    while (listNode != null) {
      System.out.print(listNode.val + " ");
      listNode = listNode.next;
    }

    System.out.println();
    System.out.println("--------------------------------");
  }
}
