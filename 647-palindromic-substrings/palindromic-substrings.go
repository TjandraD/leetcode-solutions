// 1. Initiate obj to store previously computed substr (dp)
// 2. i := 1 ... len(s)
// 3. substrStart := 0, substrEnd := i, check start == end
// 4. len(s) > 2 => Check substr between those index via dp => increment count

func countSubstrings(s string) int {
    palindromeCount := 0
    dp := map[string]bool{}

    for i := 1; i <= len(s); i++ {
        substrStart := 0
        for substrEnd := substrStart + i; substrEnd <= len(s); substrEnd++ {
            dpKey := fmt.Sprintf("%d:%d", substrStart+1, substrEnd-1)
            if isPalindrome(string(s[substrStart:substrEnd]), dp, dpKey) {
                palindromeCount++
                dpKey = fmt.Sprintf("%d:%d", substrStart, substrEnd)
                dp[dpKey] = true
            }

            substrStart++
        }
    }

    return palindromeCount
}

func isPalindrome(s string, dp map[string]bool, dpKey string) bool {
    if string(s[0]) == string(s[len(s)-1]) {
        if len(s) > 2 {
            isExist := dp[dpKey]
            if isExist {
                return true
            }
        } else {
            return true
        }
    }

    return false
}
