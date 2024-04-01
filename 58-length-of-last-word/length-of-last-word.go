func lengthOfLastWord(s string) int {
    wordsList := strings.Fields(s)
    lastWord := wordsList[len(wordsList)-1]
    return len(lastWord)
}