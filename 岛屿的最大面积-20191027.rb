


def dfs(grid, i, j, rows, columns)
  return 0 if i < 0 || j < 0 || i >= rows || j >= columns
  if grid[i][j] == 1
    grid[i][j] = 0
    result = 1 + dfs(grid, i-1, j, rows, columns) + dfs(grid, i, j-1, rows, columns) + dfs(grid, i+1, j, rows, columns) + dfs(grid, i, j+1, rows, columns)
  else
    result = 0
  end
  result
end

# @param {Integer[][]} grid
# @return {Integer}
def max_area_of_island(grid)
  rows = grid.size
  columns = grid[0].size

  max = 0
  rows.times do |i|
    columns.times do |j|
      if grid[i][j] == 1
        max = [dfs(grid, i, j, rows, columns), max].max
      end
    end
  end
  # puts max
  max
end

m = [[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]

max_area_of_island(m)