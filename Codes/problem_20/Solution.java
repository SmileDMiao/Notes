package problem_20;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

class Solution {
  public boolean isValid(String s) {

    HashMap<String, String> parentMap = new HashMap<String, String>() {
      {
        put("]", "[");
        put(")", "(");
        put("}", "{");
      }
    };

    List<String> stack = new ArrayList<String>();

    String[] str = s.split("");
    for (String i : str) {
      switch (i) {
        case "{", "[", "(":
          stack.add(i);
          System.out.println(i);
          break;
        case "}", "]", ")":
          System.out.println(i);
          if (stack.size() == 0 || !stack.get(stack.size() - 1).equals(parentMap.get(i))) {
            return false;
          } else {
            stack.remove(stack.size() - 1);
          }
          break;
      }
    }
    System.out.println(stack);

    return stack.size() == 0;
  }

  public static void main(String[] args) {

    Solution solution = new Solution();
    System.out.println(solution.isValid("()"));
  }
}
