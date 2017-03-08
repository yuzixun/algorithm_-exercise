# Write a program to find the n-th ugly number.
# Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
# For example, 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.
# Note that 1 is typically treated as an ugly number, and n does not exceed 1690.
#

# @param {Integer} n
# @return {Integer}
def nth_ugly_number(n)
  return 1 if n == 1

  result = [1]

  point_2 = 0
  point_3 = 0
  point_5 = 0

  while result.size < n
    _result_2 = result[point_2] * 2
    _result_3 = result[point_3] * 3
    _result_5 = result[point_5] * 5

    will_add_value = [_result_2, _result_3, _result_5].min
    result.push(will_add_value)

    point_2 += 1 if _result_2 == will_add_value
    point_3 += 1 if _result_3 == will_add_value
    point_5 += 1 if _result_5 == will_add_value
  end
  result.last
end
