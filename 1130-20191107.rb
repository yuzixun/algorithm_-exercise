# @param {Integer[]} arr
# @return {Integer}
def mct_from_leaf_values(arr)
  result = 0
  stack = [1<<32]

  arr.each do |item|
    while stack[-1] < item
      t = stack.pop
      result += t * [stack[-1], item].min
    end
    stack.push(item)
  end

  while stack.size > 2
    result += stack.pop * stack[-1]
  end

  result
end

arr = [6,2,4]

puts mct_from_leaf_values(arr)
puts mct_from_leaf_values([5,4,3,2,1])