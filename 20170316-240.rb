# @param {Integer[][]} matrix
# @param {Integer} target
# @return {Boolean}
#
def search_matrix(matrix, target)
  return false if matrix.empty? || matrix.first.empty?

  line_size = matrix.first.size

  columns = 0
  rows = matrix.size - 1

  while (rows >= 0) && (columns <= line_size-1)
    value = matrix[rows][columns]
    return true if value == target

    value > target ? (rows -= 1) : (columns += 1)
  end

  return false
end