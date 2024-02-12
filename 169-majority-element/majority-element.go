func majorityElement(nums []int) int {
    currElement := 0
    currCount := 0
    majorityElement := 0
    majorityCount := 0

    slices.Sort(nums)
    for _, num := range nums {
        if (currElement == num) {
            currCount++
        } else {
            currElement = num
            currCount = 1
        }

        if majorityCount < currCount {
            majorityElement = currElement
            majorityCount = currCount
        }

        if majorityCount > (len(nums) / 2) {
            break
        }
    }

    return majorityElement
}