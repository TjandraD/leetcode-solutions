# @param {Integer[][]} dominoes
# @return {Integer}
def num_equiv_domino_pairs(dominoes)
    equiv_sum = 0
    dominoes_count = {}
    dominoes.each do |domino|
        key = "#{domino[0]},#{domino[1]}"
        key = "#{domino[1]},#{domino[0]}" if domino[1] < domino[0]

        if dominoes_count[key]
            dominoes_count[key].append(domino)
        else
            dominoes_count[key] = [domino]
        end
    end

    dominoes_count.each do |_, value|
        if value.length > 1
            equiv_sum += value.combination(2).count
        end
    end

    equiv_sum
end