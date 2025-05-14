# @param {Integer[]} nums
# @return {Integer}
def remove_duplicates(nums)
    unique_items = {}
    i = 0
    while i < nums.length
        num = nums[i]
        if unique_items[num].nil?
            unique_items[num] = true
            i += 1
            next
        end

        nums.delete_at(i)
    end

    return nums.length
end