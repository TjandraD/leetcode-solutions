func tribonacci(n int) int {
    valueMap := map[int]int{}
    return findTribonacci(valueMap, n)
}

func findTribonacci(valueMap map[int]int, n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else if n == 2 {
        return 1
    } else if _, ok := valueMap[n]; ok {
        return valueMap[n]
    }

    tribonacciVal := findTribonacci(valueMap, n-1) + findTribonacci(valueMap, n-2) + findTribonacci(valueMap, n-3)
    valueMap[n] = tribonacciVal
    return tribonacciVal
}