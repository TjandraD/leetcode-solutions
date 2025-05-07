# @param {String} s
# @return {Boolean}
def is_substring_present(s)
    substring_map = {}
    substring_reversed_map = {}
    s_reversed = s.reverse

    i = 0
    while i < s.length - 1
        substring_map[s[i, 2]] = true if substring_map[s[i, 2]].nil?
        substring_reversed_map[s_reversed[i, 2]] = true if substring_reversed_map[s_reversed[i, 2]].nil?

        i += 1
    end

    substring_map.each do |keys, _|
        return true if substring_reversed_map[keys]
    end

    return false
end