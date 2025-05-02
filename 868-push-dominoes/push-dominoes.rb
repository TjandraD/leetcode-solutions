# @param {String} dominoes
# @return {String}
def push_dominoes(dominoes)
    dominoes = 'L' + dominoes + 'R'
    result = dominoes.chars

    i = 0
    while i < dominoes.length
        j = i + 1
        while j < dominoes.length && result[j] == '.'
            j += 1
        end

        process_section(result, i, j)

        i = j
    end

    result[1...-1].join
end

def process_section(dominoes, a, b)
    if dominoes[a] == dominoes[b]
        char = dominoes[a]
        (1 + a...b).each do |index|
            dominoes[index] = char
        end
    elsif dominoes[a] == 'R' && dominoes[b] == 'L'
        i = a + 1
        j = b - 1
        while i < j
            dominoes[i] = 'R'
            dominoes[j] = 'L'
            i += 1
            j -= 1
        end
    end
end