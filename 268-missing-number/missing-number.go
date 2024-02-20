// 1. Sort the array asc
// 2. Check each item, verify that nums[i] = i
func missingNumber(nums []int) int {
    slices.Sort(nums)
    for i := 0; i < len(nums); i++ {
        if nums[i] != i {
            return i
        }
    }
    return len(nums);
}