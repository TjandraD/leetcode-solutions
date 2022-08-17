class Solution {
    public int rob(int[] nums) {
        int rob = 0;
        int notRob = 0;
        for (int i : nums) {
            int curRob = notRob + i;
            notRob = Math.max(notRob, rob);
            rob = curRob;
        }
        return Math.max(rob, notRob);
    }
}