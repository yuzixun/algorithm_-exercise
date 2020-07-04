# @param {Integer[][]} dungeon
# @return {Integer}
def calculate_minimum_hp(dungeon)
  return 0 if dungeon.empty? || dungeon[0].empty?

  x = dungeon.size
  y = dungeon[0].size

  current_x = x - 1
  current_y = y - 1

  while current_y >= 0
    while current_x >= 0
      if (current_x == x - 1) && (current_y == y - 1)
        dungeon[current_x][current_y] = [1, 1 - dungeon[current_x][current_y]].max
      elsif current_x == x - 1
        dungeon[current_x][current_y] = [1, dungeon[current_x][current_y+1] - dungeon[current_x][current_y]].max
      elsif current_y == y - 1
        dungeon[current_x][current_y] = [1, dungeon[current_x+1][current_y] - dungeon[current_x][current_y]].max
      else
        dungeon[current_x][current_y] = [1, [dungeon[current_x][current_y+1], dungeon[current_x+1][current_y]].min - dungeon[current_x][current_y]].max
      end
      current_x -= 1
    end
    current_y -= 1; current_x = x - 1
  end

  dungeon[0][0]
end
