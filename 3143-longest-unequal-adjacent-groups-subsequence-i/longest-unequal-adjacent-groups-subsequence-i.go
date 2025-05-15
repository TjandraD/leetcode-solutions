func getLongestSubsequence(words []string, groups []int) []string {
    if len(words) == 1 {
        return words
    }

    subsequence := []string{}
    for i, word := range words {
        if i == 0 {
            subsequence = append(subsequence, word)
            continue
        }

        if groups[i] != groups[i-1] {
            subsequence = append(subsequence, word)
        }
    }

    return subsequence
}