func findFinalValue(nums []int, original int) int {
    numMap := map[int]bool{}
    for _, num := range nums {
        numMap[num] = true
    }

    if numMap[original] == false {
        return original
    }

    final := original
    for {
        final *= 2
        if numMap[final] == true {
            continue
        }

        break
    }

    return final
}