# @param {Integer} k
# @param {Integer[]} prices
# @return {Integer}
def max_profit(k, prices)
  return 0 if prices.empty? || k.zero?
  min_int = -1<<32
  profits = []
  prices.size.times do |i|
    profits[i] = (k+1).times.map { [min_int, min_int] }
  end

  profits[0][0] = [0, -prices[0]]

  for i in (1...prices.size) do
    for kk in (0..k) do
      if kk == 0
        profits[i][0][0] = profits[i-1][0][0]
        profits[i][0][1] = [ profits[i-1][0][1], profits[i-1][0][0] - prices[i] ].max
      else
        profits[i][kk][0] = [ profits[i-1][kk][0], profits[i-1][kk-1][1] + prices[i] ].max
        profits[i][kk][1] = [ profits[i-1][kk][1], profits[i-1][kk][0] - prices[i] ].max
      end
    end
  end

  profits[-1].map(&:first).max
end

puts max_profit(2, [2,4,1])
puts max_profit(2, [3,2,6,5,0,3])