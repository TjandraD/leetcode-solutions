# @param {Integer[]} nums
# @return {Integer[]}
def build_array(nums)
    nums_map = {}
    i = 0
    while i < nums.length do
        nums_map[i] = nums[i]
        i += 1
    end

    ans = []
    i = 0
    while i < nums.length do
        ans[i] = nums_map[nums_map[i]]
        i += 1
    end

    ans
end