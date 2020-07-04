# @param {Integer[]} satisfaction
# @return {Integer}
def max_satisfaction(satisfaction)
  satisfaction = satisfaction.sort { |a, b| b <=> a }
  return 0 if satisfaction.first < 0

  result = []
  cache = []

  satisfaction.each_with_index do |i, index|
    if index == 0
      result[index] = i
      cache[index] = i
    else
      result[index] = cache[index-1] + result[index-1]+ i
      cache[index] = cache[index-1] +i
    end

  end

  result.max
end

p max_satisfaction([-1,-8,0,5,-9])
p max_satisfaction([4,3,2])
p max_satisfaction([-2,5,-1,0,3,-3])