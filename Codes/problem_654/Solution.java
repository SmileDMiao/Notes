package problem_654;

import java.util.stream.IntStream;

import structures.TreeNode;

public class Solution {
  public TreeNode constructMaximumBinaryTree(int[] nums) {
    if (nums.length == 0) {
      return null;
    }

    Integer max = -1;
    Integer index = -1;

    for (Integer i = 0; i < nums.length; i++) {
      if (nums[i] >= max) {
        max = nums[i];
        index = i;
      }
    }

    TreeNode root = new TreeNode(max, null, null);
    int[] left = IntStream.range(0, index)
        .map(i -> nums[i])
        .toArray();

    int[] right = IntStream.range(index+1, nums.length)
        .map(i -> nums[i])
        .toArray();

    root.left = constructMaximumBinaryTree(left);
    root.right = constructMaximumBinaryTree(right);

    return root;
  }

  public static void main(String[] args) {
    int[] nums = { 3, 2, 1, 6, 0, 5 };
    Solution solution = new Solution();
    solution.constructMaximumBinaryTree(nums);
  }
}
