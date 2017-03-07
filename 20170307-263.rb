# Write a program to check whether a given number is an ugly number.
# Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 6, 8 are ugly while 14 is not ugly since it includes another prime factor 7.
# Note that 1 is typically treated as an ugly number.

# @param {Integer} num
# @return {Boolean}
def is_ugly(num)
  return false if num == 0
  return true if num == 1

  case 0
  when num % 2
    is_ugly(num/2)
  when num % 3
    is_ugly(num/3)
  when num % 5
    is_ugly(num/5)
  else
    return false
  end
end

puts is_ugly(15)