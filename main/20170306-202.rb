# Write an algorithm to determine if a number is "happy".
# A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
# Example: 19 is a happy number
# 1^2 + 9^2 = 82
# 8^2 + 2^2 = 68
# 6^2 + 8^2 = 100
# 1^2 + 0^2 + 0^2 = 1

# author: yuzixun
# time: O(lgn)
# space: O(lgn)
# @param {Integer} n
# @return {Boolean}
def get_result_by(string)
  string.chars.reduce(0) do |sum, element|
    sum += element.to_i**2
  end
end

def is_happy(n)
  sums = []

  loop do
    n = get_result_by(n.to_s)

    return true if n == 1
    return false if sums.include?(n)

    sums.push(n)
  end
end

puts is_happy(19)