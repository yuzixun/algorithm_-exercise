# @param {Integer[][]} matrix
# @param {Integer} target
# @return {Boolean}
#
def search_matrix(matrix, target)
  return false if matrix.size == 0

  columns = matrix[0].size
  rows = matrix.size

  left = 0
  right = columns*rows - 1

  while left <= right
    mid = (left+right)/2
    current = matrix[mid/columns][mid%columns]

    case true
    when current == target
      return true
    when current < target
      left = mid + 1
    when current > target
      right = mid - 1
    end
  end

  return false
end
