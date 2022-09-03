package problem_203;

import java.util.Arrays;
import java.util.List;

import structures.ListNode;

class Solution {
  public ListNode removeElements(ListNode head, int val) {

    ListNode dummy = new ListNode(-1);
    dummy.next = head;

    var cursor = dummy;

    while (cursor != null && cursor.next != null) {
      if (cursor.next.val == val) {
        cursor.next = cursor.next.next;
      } else {
        cursor = cursor.next;
      }
    }

    return dummy.next;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();

    List<Integer> values = Arrays.asList(7, 7, 7, 7);

    ListNode head = ListNode.createListNode(values);
    ListNode.print(head);

    ListNode result = solution.removeElements(head, 7);

    ListNode.print(result);
  }
}
