func reversePrefix(word string, ch byte) string {
    chIndex := 0
    charList := strings.Split(word, "")
    charMap := map[int]string{}

    for i, charItem := range charList {
        charMap[i] = charItem
        if charItem == string(ch) && chIndex == 0 {
            if i == 0 {
                break
            }
            chIndex = i
        }
    }

    if chIndex == 0 {
        return word
    }

    reversedWord := ""
    for i := 0; i <= chIndex; i++ {
        reversedWord += charMap[chIndex-i]
    }

    for i := chIndex+1; i < len(word); i++ {
        reversedWord += charMap[i]
    }

    return reversedWord
}