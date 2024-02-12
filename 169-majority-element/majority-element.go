func majorityElement(nums []int) int {
    numCount := map[int]int{}
    majorityElement := 0
    for _, num := range nums {
        count := numCount[num]
        count++
        numCount[num] = count

        if count > numCount[majorityElement] {
            majorityElement = num
        }

        if numCount[majorityElement] > (len(nums) / 2) {
            break
        }
    }

    return majorityElement
}