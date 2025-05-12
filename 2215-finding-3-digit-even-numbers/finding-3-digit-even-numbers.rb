# @param {Integer[]} digits
# @return {Integer[]}
def find_even_numbers(digits)
    digits_map = {}
    digits.each do |digit|
        digits_map[digit] = 0 if digits_map[digit].nil?
        digits_map[digit] += 1
    end

    nums = []
    i = 100
    while i < 999 do
        map_copy = digits_map.clone
        current_num = i.to_s

        digits_exist = true
        current_num.chars.each do |digit|
            if map_copy[digit.to_i].nil? || map_copy[digit.to_i].zero?
                digits_exist = false
                break
            end

            map_copy[digit.to_i] -= 1
        end

        nums.push(i) if digits_exist
        i += 2
    end

    nums
end