# For an integer n, we call k>=2 a good base of n, if all digits of n base k are 1.
# Now given a string representing n, you should return the smallest good base of n in string format.
#
# Example 1:
# Input: "13"
# Output: "3"
# Explanation: 13 base 3 is 111.
#
# Example 2:
# Input: "4681"
# Output: "8"
# Explanation: 4681 base 8 is 11111.
#
# Example 3:
# Input: "1000000000000000000"
# Output: "999999999999999999"
# Explanation: 1000000000000000000 base 999999999999999999 is 11.
#
# Note:
# The range of n is [3, 10^18].
# The string representing n is always valid and will not have leading zeros.
#

# n = k^m + k^(m-1) + ... + k + 1
# => n-1 = k^m + k^(m-1) + ... + k
# => n-1 = k (k^(m-1) + k^(m-2) + ... + k + 1) ...... [1]

# n = k^m + k^(m-1) + ... + k + 1
# => n-k^m = k^(m-1) + k^(m-2) + ... + k + 1 ...... [2]
# [1] - [2]
# => (k^(m+1) - 1)/(k - 1) = n .... [3]
#
# n = k^m + k^(m-1) + ... + k + 1
# => n > k^m
# => m-th root of n > k .... [4]
# => k+1 > m-th root of n > k. .... from [4] and [5]
# Thus ⌊m-th root of n⌋ is the only candidate that needs to be tested. [6]
#
# We also know that the smallest base is 2 so we can find our m must be between 2 and log2n else m is (n-1) [7]

# @param {String} n
# @return {String}
def smallest_good_base(n)
  n = n.to_i

  m = Math.log(n, 2).to_i

  while m > 1
    k = (n**(1.0/m)).to_i

    return k.to_s if (k**(m+1)-1)/(k-1) == n

    m -= 1
  end
  return (n - 1).to_s
end
