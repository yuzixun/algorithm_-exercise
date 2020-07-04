
# @param {Integer} n
# @return {Integer}
def count_largest_group(n)
  group = Hash.new(0)
  1.upto(n).each do |i|
    cur = 0

    while i != 0
      cur += (i % 10)
      i /= 10
    end
    group[cur] += 1
  end
  m = group.values.max
  group.find_all { |_, v| v == m }.size
end


p count_largest_group(13)
p count_largest_group(2)
p count_largest_group(15)
p count_largest_group(24)
p count_largest_group(10**4)