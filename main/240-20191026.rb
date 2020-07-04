# @param {Integer[][]} matrix
# @param {Integer} target
# @return {Boolean}
def search_matrix(matrix, target)
  return false if matrix.empty? || matrix[0].empty?

  r_index = 0
  c_index = matrix[0].size-1

  while r_index <= matrix.size-1 && c_index >= 0
    puts "rindex is #{r_index}, cindex is #{c_index}"
    current = matrix[r_index][c_index]
    case true
    when current == target
      return true
    when current < target
      r_index+=1
    when current > target
      c_index -=1
    end
  end

  false
end


# matrix = [
#   [1,   4,  7, 11, 15],
#   [2,   5,  8, 12, 19],
#   [3,   6,  9, 16, 22],
#   [10, 13, 14, 17, 24],
#   [18, 21, 23, 26, 30]
# ]

# puts search_matrix(matrix, 20)

matrix = [[1,1]]

puts search_matrix(matrix, 2)