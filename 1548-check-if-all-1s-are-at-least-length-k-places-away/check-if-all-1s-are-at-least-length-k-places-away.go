func kLengthApart(nums []int, k int) bool {
    lastDistance := 1
    found := false
    for _, i := range nums {
        if i != 1 && !found {
            continue
        } else if i == 1 && !found {
            found = true
            lastDistance = 0
            continue
        }

        if i == 1 {
            if lastDistance < k {
                return false
            }

            lastDistance = 0
            continue
        }

        lastDistance += 1
    }

    return true
}