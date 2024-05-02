func findMaxK(nums []int) int {
    slices.Sort(nums)
    numbersMap := map[int]bool{}
    biggestNum := -1

    for _, num := range nums {
        if num < 0 {
            numbersMap[num] = true
        } else {
            negativeNum := num * -1
            if numbersMap[negativeNum] {
                biggestNum = num
            }
        }
    }

    return biggestNum
}