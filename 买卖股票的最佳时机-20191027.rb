# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
  i = 0
  size = prices.size
  max = 0
  prev_min = prices[0]

  while i <= size-1
    prev_min = [prev_min, prices[i]].min
    max = [max, (prices[i]-prev_min)].max
    i+=1
  end

  max
end


puts max_profit([7,1,5,3,6,4])
puts max_profit([7,6,4,3,1])
puts max_profit([1,2])