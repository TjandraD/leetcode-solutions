func findJudge(n int, trust [][]int) int {
    if n == len(trust) {
        return -1
    }

    if n == 1 && len(trust) == 0 {
        return 1
    }

    trustMap := map[int][]int{}
    trustCount := map[int]int{}
    for _, trustees := range trust {
        truster := trustees[0]
        trustee := trustees[1]

        if _, ok := trustMap[trustee]; !ok {
            trustMap[trustee] = []int{}
        }
        trustMap[trustee] = append(trustMap[trustee], truster)

        if _, ok := trustCount[truster]; !ok {
            trustCount[truster] = 0
        }
        trustCount[truster]++
    }

    for i := 1; i <= n; i++ {
        if _, ok := trustMap[i]; ok {
            _, countExist := trustCount[i]
            if len(trustMap[i]) == (n-1) && !countExist {
                return i
            }
        }
    }

    return -1
}