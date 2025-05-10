# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Integer}
def min_sum(nums1, nums2)
    sum_nums1 = nums1.sum
    sum_nums2 = nums2.sum
    nums1_zeroes = nums1.count(0)
    nums2_zeroes = nums2.count(0)

    final_sum_nums1 = sum_nums1 + nums1_zeroes
    final_sum_nums2 = sum_nums2 + nums2_zeroes

    return -1 if (final_sum_nums1 > sum_nums2 && nums2_zeroes == 0) || (final_sum_nums2 > sum_nums1 && nums1_zeroes == 0)

    return final_sum_nums1 if final_sum_nums1 > final_sum_nums2

    return final_sum_nums2
end