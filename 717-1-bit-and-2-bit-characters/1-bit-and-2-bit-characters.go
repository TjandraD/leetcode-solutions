func isOneBitCharacter(bits []int) bool {
    i := 0
    for i < len(bits) {
        if bits[i] == 1 && i < len(bits)-1 {
            if i+2 == len(bits) {
                return false
            }

            i += 2
            continue
        }

        if i == len(bits)-1 {
            if bits[i] == 0 {
                break
            }

            return false
        }

        i++
    }

    return true
}