# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
	return 0 if prices.empty?
	min_int = -2<<32
	profits = []

	prices.size.times do |i|
		profits[i] = [[min_int, min_int], [min_int, min_int], [min_int, min_int]]
	end

	profits[0][0] = [0, -prices[0]]

	# i: 天数， j 交易次数， k 是否有股票
	for i in (1...prices.size) do
		# puts "profits is #{profits[i-1]}, i is #{i}"
		profits[i][0][0] = profits[i-1][0][0]
		profits[i][0][1] = [ profits[i-1][0][1], profits[i-1][0][0] - prices[i] ].max

		profits[i][1][0] = [ profits[i-1][1][0], profits[i-1][0][1] + prices[i] ].max
		profits[i][1][1] = [ profits[i-1][1][1], profits[i-1][1][0] - prices[i] ].max

		profits[i][2][0] = [ profits[i-1][2][0], profits[i-1][1][1] + prices[i] ].max
	end

	profits[-1].map(&:first).max
end

# prices = [3,3,5,0,0,3,1,4]
prices = []
puts max_profit(prices)