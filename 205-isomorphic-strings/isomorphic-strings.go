func isIsomorphic(s string, t string) bool {
    charMap := map[byte]byte{}
    charUsed := map[byte]bool{}

    for i := 0; i < len(s); i++ {
        sChar := s[i]
        tChar := t[i]

        if _, isExist := charMap[sChar]; !isExist {
            if _, isUsed := charUsed[tChar]; isUsed {
                return false
            }

            charMap[sChar] = tChar
            charUsed[tChar] = true
            continue
        }

        storedChar := charMap[sChar]
        if storedChar != tChar {
            return false
        }
    }

    return true
}