# @param {Integer[][]} matrix
# @param {Integer} target
# @return {Boolean}
#

def search_matrix(matrix, target)
  return false if matrix.empty? || matrix.first.empty?

  matrix_size = matrix.size

  columns = matrix[0].size - 1
  rows = 0

  while (rows <= matrix_size-1) && (columns >= 0)
    case true
    when matrix[rows][columns] == target
      return true
    when matrix[rows][columns] > target
      columns -= 1
    when matrix[rows][columns] < target
      rows += 1
    end
  end

  return false
end


matrix = [
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]

target1 = 5
target2 = 20
puts search_matrix(matrix, target1)
puts search_matrix(matrix, target2)