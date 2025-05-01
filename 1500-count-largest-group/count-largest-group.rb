# @param {Integer} n
# @return {Integer}
def count_largest_group(n)
    groups = {}
    i = 1

    while i <= n
        sum = digit_sum(i)

        if groups[sum]
            groups[sum] += 1
        else
            groups[sum] = 1
        end

        i += 1
    end

    largest_size = 0
    group_sizes = {}

    groups.each do |_,val|
        next if val < largest_size

        if group_sizes[val]
            group_sizes[val] += 1
        else
            group_sizes[val] = 1
        end

        largest_size = val if val > largest_size
    end

    group_sizes[largest_size]
end

def digit_sum(n)
    return 0 if n == 0
    return n % 10 + digit_sum(n/10)
end