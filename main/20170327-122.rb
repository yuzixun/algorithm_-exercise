# Say you have an array for which the ith element is the price of a given stock on day i.
# Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times). However, you may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
# Subscribe to see which companies asked this question.

# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
  sum = 0
  prices.size.times do |index|
    sum += [prices.fetch(index+1, 0) - prices[index], 0].max
  end
  sum
end

puts max_profit([1])