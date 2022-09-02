package problem_70;

import java.util.ArrayList;

public class Solution {

  public int climbStairs(int n) {
    if (n <=3){
      return n;
    }

    ArrayList<Integer> dp = new ArrayList<>();
    dp.add(1);
    dp.add(2);

    for(int i = 2; i < n; i++) {
      dp.add(dp.get(i-1) + dp.get(i-2));
    }

    System.out.println(dp);

    return dp.get(dp.size() - 1);
  }

  public static void main(String[] args) {
    Solution solution = new Solution();

    int steps = 10;
    System.out.println(solution.climbStairs(steps));
  }

}
