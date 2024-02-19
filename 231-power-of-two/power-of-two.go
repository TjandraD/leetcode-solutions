func isPowerOfTwo(n int) bool {
    for i := 0; math.Pow(float64(2), float64(i)) <= float64(n); i++ {
        if math.Pow(float64(2), float64(i)) == float64(n) {
            return true
        }
    }

    return false
}