# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    cached_nums = {}
    i = 0
    while i < nums.length()
        cached_nums[nums[i]] = i if cached_nums[nums[i]].nil?
        i += 1
    end

    i = 0
    while i < nums.length()
        diff = target - nums[i]
        return [i, cached_nums[diff]] if !cached_nums[diff].nil? && cached_nums[diff] != i
        i += 1
    end

    return [0, 0]
end
