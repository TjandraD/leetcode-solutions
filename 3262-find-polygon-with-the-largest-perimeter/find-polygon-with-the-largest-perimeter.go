// 1. Sort the array with asc order
// 2. Loop through the nums, sum the total
// 3. Loop through nums from the last to the third items
// 4. If currItem < sumSides, return currItem

func largestPerimeter(nums []int) int64 {
    slices.Sort(nums)

    var perimeter int64
    for _, num := range nums {
        perimeter += int64(num)
    }

    for i := len(nums)-1; i >= 2; i-- {
        sumSides := perimeter - int64(nums[i])
        if sumSides > int64(nums[i]) {
            return perimeter
        }

        perimeter = sumSides
    }

    return -1
}