package problem_35;

import java.util.Arrays;

public class Solution {
  public int searchInsert(int[] nums, int target) {
    Integer left = 0;
    Integer right = nums.length - 1;

    while (left <= right) {
      Integer middle = left + (right - left) / 2;
      if (nums[middle] > target) {
        right = middle - 1;
      } else if (nums[middle] < target) {
        left = middle + 1;
      } else {
        return middle;
      }
    }

    return right + 1;
  }

  public int searchInsertWithApi(int[] nums, int target) {
    int i = Arrays.binarySearch(nums, target);
    return i >= 0 ? i : -i - 1;
  }

  public static void main(String[] args) {
    Solution solution = new Solution();

    int[] nums = { 1, 3, 5, 6 };
    System.out.println(solution.searchInsert(nums, 0));
  }
}
