# @param {Integer[]} nums
# @return {Integer}
def count_subarrays(nums)
    condition_met = 0
    i = 2
    while i < nums.length()
        if ((nums[i-2] + nums[i]) * 2) == nums[i-1]
            condition_met += 1
        end

        i += 1
    end

    condition_met
end