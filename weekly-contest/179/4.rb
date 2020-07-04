# @param {Integer} n
# @param {Integer[][]} edges
# @param {Integer} t
# @param {Integer} target
# @return {Float}
def frog_position(n, edges, t, target)
  from_to = {}
  count = Hash.new(0)

  edges.each do |edge|
    edge.sort!

    from_to[edge[1]] = edge[0]
    count[edge[0]] +=1
  end

  result = 1.0
  cur = target
  counter = 0

  while from_to[cur]
    result /= count[from_to[cur]]
    cur = from_to[cur]
    counter+=1
  end

  # puts "counter is #{counter}, t is #{t}, result is #{result}, count is #{count}, target is #{target}, fromto is #{from_to}"
  case true
  when counter == t
    return result
  when counter < t
    return count[target] == 0 ? result : 0
  end
  return 0
end
n = 7
edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]
t = 20
target = 6
puts frog_position(n, edges, t, target)



n = 7
edges = [[2,1],[3,1],[4,2],[5,4],[6,5],[7,2]]
t = 3
target = 2
puts frog_position(n, edges, t, target)

n = 8
edges =  [[2,1],[3,2],[4,1],[5,1],[6,4],[7,1],[8,7]]
t =  7
target = 7
puts frog_position(n, edges, t, target)

n = 10
edges =  [[2,1],[3,2],[4,2],[5,2],[6,5],[7,1],[8,3],[9,1],[10,1]]
t =  1
target =  9
puts frog_position(n, edges, t, target)


n = 7
edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]
t = 2
target = 4
puts frog_position(n, edges, t, target)

n = 7
edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]]
t = 1
target = 7
puts frog_position(n, edges, t, target)

