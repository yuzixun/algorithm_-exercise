# Write a program to find the nth super ugly number.

# Super ugly numbers are positive numbers whose all prime factors are in the given prime list primes of size k. For example, [1, 2, 4, 7, 8, 13, 14, 16, 19, 26, 28, 32] is the sequence of the first 12 super ugly numbers given primes = [2, 7, 13, 19] of size 4.

# Note:
# (1) 1 is a super ugly number for any given primes.
# (2) The given numbers in primes are in ascending order.
# (3) 0 < k ≤ 100, 0 < n ≤ 106, 0 < primes[i] < 1000.
# (4) The nth super ugly number is guaranteed to fit in a 32-bit signed integer.
#

# @param {Integer} n
# @param {Integer[]} primes
# @return {Integer}

def nth_super_ugly_number(n, primes)
  return 1 if n == 1

  result = [1]
  points = [0]*primes.size

  while result.size < n
    _temp_result = []
    primes.each_with_index do |prime, index|
      _temp_result.push(result[points[index]] * prime)
    end

    will_add_value = _temp_result.min
    result.push(will_add_value)

    points.each_with_index do |prime, index|
      points[index] += 1 if will_add_value == _temp_result[index]
    end
  end

  result.last
end

puts nth_super_ugly_number 100000, [7,19,29,37,41,47,53,59,61,79,83,89,101,103,109,127,131,137,139,157,167,179,181,199,211,229,233,239,241,251]
