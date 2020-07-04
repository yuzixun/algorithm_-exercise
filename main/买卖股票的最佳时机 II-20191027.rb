# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
  profit = 0
  size = prices.size

  i = 1
  while i < size
    profit += [0, prices[i] - prices[i-1]].max
    i+=1
  end

  profit
end

puts max_profit([7,1,5,3,6,4])
puts max_profit([1,2,3,4,5])
puts max_profit([7,6,4,3,1])