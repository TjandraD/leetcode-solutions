func sortColors(nums []int)  {
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            smallestIndex := findSmallestIndex(nums, j)
            if smallestIndex != j {
                nums[smallestIndex] += nums[j]
                nums[j] = nums[smallestIndex] - nums[j]
                nums[smallestIndex] -= nums[j]
            }
        }
    }
}

func findSmallestIndex(nums []int, limit int) int {
    for i := 0; i < limit; i++ {
        if nums[i] > nums[limit] {
            return i
        }
    }

    return limit
}