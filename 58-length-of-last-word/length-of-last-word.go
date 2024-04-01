func lengthOfLastWord(s string) int {
    prevChar := ' '
    len := 0
    for _, stringChar := range s {
        if prevChar == ' ' && stringChar != ' ' {
            len = 1
        } else if stringChar != ' ' {
            len++
        }

        prevChar = stringChar
    }

    return len
}