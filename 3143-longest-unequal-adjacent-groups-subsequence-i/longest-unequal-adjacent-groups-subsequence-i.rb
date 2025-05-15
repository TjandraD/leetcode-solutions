# @param {String[]} words
# @param {Integer[]} groups
# @return {String[]}
def get_longest_subsequence(words, groups)
    return words if words.length == 1

    subsequence = []
    words.each_with_index do |word, i|
        if i == 0
            subsequence.push(word)
            next
        end

        subsequence.push(word) if groups[i] != groups[i - 1]
    end

    return subsequence
end