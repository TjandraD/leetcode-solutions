class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> numsMap = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int leftToTarget = target - nums[i];
            if (numsMap.get(leftToTarget) != null) {
                return new int[]{numsMap.get(leftToTarget), i};
            }
            numsMap.put(nums[i], i);
        }
        
        return new int[]{0, 0};
    }
}