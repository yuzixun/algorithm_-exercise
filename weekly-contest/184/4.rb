# @param {Integer} n
# @return {Integer}
def num_of_ways(n)
  same, diff = 6, 6
  mod = 10**9 + 7
  (n-1).times do |i|
    n_s = 3 * same + 2 * diff
    n_d = 2 * same + 2 * diff

    same, diff = n_s%mod, n_d%mod
  end

  (same + diff)%mod
end

# p num_of_ways(1)
# p num_of_ways(2)
# p num_of_ways(3)
p num_of_ways(5000)