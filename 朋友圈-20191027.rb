# def clear_related(m, i, j, max_i, max_j)
#   direction = [[-1, 0], [0, -1], [1, 0], [0, 1]]

# end

# # @param {Integer[][]} m
# # @return {Integer}
# def find_circle_num(m)
#   i = m.size
#   j = m[0].size
#   count = 0
#   i.times do |ii|
#     j.times do |jj|
#       next if m[ii][jj] == 0
#       count += 1
#       clear_related(m, ii, jj, i-1, j-1)
#     end
#   end

#   count
# end


def find_root(circles, node)
  puts "circles is #{circles}, node is #{node}"
  while circles[node] != node
    node = circles[node]
  end

  node
end
# @param {Integer[][]} m
# @return {Integer}
def find_circle_num(m)
  return if m.empty?

  i = m.size
  j = m[0].size

  # 初始化
  circles = {}
  i.times do |ri|
    circles[ri] = ri
  end

  i.times do |ii|
    j.times do |jj|
      next if ii == jj

      # 两者有交集，必然指向同一个 root，如果没有，则需要更新
      if m[ii][jj] == 1 && find_root(circles, ii) != find_root(circles, jj)
        circles[find_root(circles, jj)] = find_root(circles, ii)
      end
    end
  end

  r = 0
  circles.each do |k, v|
    r += 1 if k == v
  end
  r
end


m = [[1,1,0],
 [1,1,0],
 [0,0,1]]
m = [[1,1,0],
 [1,1,1],
 [0,1,1]]

# m = [[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]
puts find_circle_num(m)

