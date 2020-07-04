# @param {Integer[]} g
# @param {Integer[]} s
# @return {Integer}
def find_content_children(g, s)
  g.sort!
  s.sort!

  result = 0

  g_size = g.size
  g_index = 0
  s.size.times do |s_index|
    if (g_index <= g_size-1) && (s[s_index] >= g[g_index])
      result += 1
      g_index += 1
    end
  end

  result
end


puts find_content_children([1,2], [1,2,3])
puts find_content_children([1,2,3], [1, 1])