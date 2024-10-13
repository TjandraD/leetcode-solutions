# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    i = 0
    until i == nums.length() do
        j = i + 1
        until j == nums.length() do
            return [i, j] if nums[i] + nums[j] == target
            j += 1
        end
        i += 1
    end

    return [0, 0]
end
