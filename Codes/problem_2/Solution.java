package problem_2;

import java.util.Arrays;
import java.util.List;

import strustures.ListNode;

class Solution {
  public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
    ListNode result = new ListNode(-1);
    ListNode head = result;

    int carry = 0;

    while (l1 != null || l2 != null) {
      int l1Val = l1 != null ? l1.val : 0;
      int l2Val = l2 != null ? l2.val : 0;

      int sum = l1Val + l2Val + carry;

      carry = sum >= 10 ? 1 : 0;
      sum = sum % 10;

      result.next = new ListNode(sum);
      result = result.next;
      l1 = l1 != null ? l1.next : null;
      l2 = l2 != null ? l2.next : null;
    }

    if (carry != 0) {
      result.next = new ListNode(carry);
    }

    return head.next;
  }

  public static void main(String[] args) {
    List<Integer> values1 = Arrays.asList(1,2,3,4,5);
    ListNode l1 = ListNode.createListNode(values1);

    List<Integer> values2 = Arrays.asList(5,4,3);
    ListNode l2 = ListNode.createListNode(values2);

    Solution solution = new Solution();
    ListNode.print(solution.addTwoNumbers(l1, l2));
  }

}
