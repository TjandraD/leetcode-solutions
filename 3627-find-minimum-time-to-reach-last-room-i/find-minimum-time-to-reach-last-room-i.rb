# @param {Integer[][]} move_time
# @return {Integer}
def min_time_to_reach(move_time)
    m = move_time.length
    n = move_time[0].length
    target = "#{m-1}:#{n-1}"

    move_time_map = {}
    costs = {}
    move_time.each_with_index do |row, i|
        move_time_map[i] = {}
        row.each_with_index do |room, j|
            move_time_map[i][j] = room
            costs["#{i}:#{j}"] = Float::INFINITY
        end
    end

    costs["0:0"] = 0
    visits = {}
    queue = [[0, 0, 0]]
    movements = [[-1, 0], [0, -1], [1, 0], [0, 1]]

    while queue.length > 0 do
        cost, i, j = queue.shift
        room = "#{i}:#{j}"

        break if room == target
        next if visits[room]
        visits[room] = true

        movements.each do |movement|
            new_i = i + movement[0]
            new_j = j + movement[1]

            if (new_i >= 0 && new_i < m) && (new_j >= 0 && new_j < n)
                adj_room = "#{new_i}:#{new_j}"
                if i == 0 && j == 0
                    costs[adj_room] = move_time_map[new_i][new_j] + 1
                    queue.push([costs[adj_room], new_i, new_j])
                    next
                end

                adj_room_cost = [cost + 1, move_time_map[new_i][new_j] + 1].max
                costs[adj_room] = adj_room_cost if adj_room_cost < costs[adj_room]
                queue = queue_push(queue, [costs[adj_room], new_i, new_j])
            end
        end
    end

    return costs[target]
end

def queue_push(queue, item)
    if queue.empty?
        return [item]
    end

    i = 0
    while i < queue.length do
        break if queue[i][0] > item[0]
        i += 1
    end

    return queue.insert(i, item)
end