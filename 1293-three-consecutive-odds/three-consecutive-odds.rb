# @param {Integer[]} arr
# @return {Boolean}
def three_consecutive_odds(arr)
    first_odd_i = -1
    arr.each_with_index do |num, i|
        first_odd_i = -1 if num.even? && first_odd_i >= 0
        first_odd_i = i if num.odd? && first_odd_i.negative?

        return true if first_odd_i == i - 2 && first_odd_i >= 0
    end

    return false
end