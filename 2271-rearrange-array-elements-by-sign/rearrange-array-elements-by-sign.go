func rearrangeArray(nums []int) []int {
    positiveNums := []int{}
    negativeNums := []int{}
    sortedArray := []int{}

    for _, num := range nums {
        if num > 0 {
            positiveNums = append(positiveNums, num)
            continue
        }
        negativeNums = append(negativeNums, num)
    }

    for i := 0; i < len(positiveNums); i++ {
        sortedArray = append(sortedArray, positiveNums[i], negativeNums[i])
    }

    return sortedArray
}