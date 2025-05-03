# @param {Integer[]} tops
# @param {Integer[]} bottoms
# @return {Integer}
def min_domino_rotations(tops, bottoms)
    tile_counts = {}
    i = 0

    while i < tops.length do
        top_tile = tops[i]
        bottom_tile = bottoms[i]

        if tile_counts[top_tile]
            tile_counts[top_tile] += 1
        else
            tile_counts[top_tile] = 1
        end

        if tile_counts[bottom_tile]
            tile_counts[bottom_tile] += 1
        else
            tile_counts[bottom_tile] = 1
        end

        if top_tile == bottom_tile
            # only count 1 per domino
            tile_counts[bottom_tile] -= 1
        end

        i += 1
    end

    same_value = 0
    (1..6).each do |tile|
        same_value = tile if tile_counts[tile] == tops.length
    end

    return -1 if same_value == 0

    top_rotation = 0
    bottom_rotation = 0
    i = 0
    while i < tops.length do
        top_tile = tops[i]
        bottom_tile = bottoms[i]

        top_rotation += 1 if top_tile != same_value
        bottom_rotation += 1 if bottom_tile != same_value

        i += 1
    end

    return top_rotation if top_rotation < bottom_rotation
    bottom_rotation
end