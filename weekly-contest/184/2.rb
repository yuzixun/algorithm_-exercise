# @param {Integer[]} queries
# @param {Integer} m
# @return {Integer[]}
def process_queries(queries, m)
  arr = (1..m).to_a

  result = []

  queries.each do |query|
    ri = arr.find_index { |num| num == query }
    result.push(ri)
    arr = [query] + arr[0...ri] + arr[ri+1..-1]

  end

  result
end

p process_queries([3,1,2,1],5)
p process_queries([4,1,2,2],4)
p process_queries([7,5,5,8,3],8)