# @param {Integer} num
# @return {Integer}
def maximum69_number (num)
    output = ""
    num.to_s.chars.each_with_index do |num_char, i|
        output += num_char if num_char == "9"

        if num_char == "6"
            output += "9"
            output += num.to_s.chars[i+1, num.to_s.chars.length-1].join("")
            break
        end
    end

    output.to_i
end