# @param {Character[][]} board
# @param {String} word
# @return {Boolean}
#
def match?(y, x, board, word)
  return true if word == ''
  return false if x < 0 || y < 0 || y == board.size || x == board[0].size || board[y][x] != word[0]

  board[y][x] = '*'

  surplus_word = word[1..-1]
  result = match?(y-1, x, board, surplus_word) || match?(y+1, x, board, surplus_word) || match?(y, x-1, board, surplus_word) || match?(y, x+1, board, surplus_word)

  board[y][x] = word[0]
  return result
end

def exist(board, word)

  x = board[0].size
  y = board.size

  y.times { |i| x.times { |j| return true if match?(i, j, board, word) } }

  return false
end


board = [
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

puts exist(board, 'ABCCED')
puts exist(board, 'SEE')
puts exist(board, 'ABCB')
