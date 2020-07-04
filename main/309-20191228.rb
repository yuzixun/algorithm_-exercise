# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
	dp_0 = 0
	dp_1 = -1 << 32
	dp_pre_0 = 0
	prices.each_with_index do |price, i|
		temp = dp_0

		dp_0 = [dp_0, dp_1 + price].max
		dp_1 = [dp_1, dp_pre_0 - price].max

		dp_pre_0 = temp
	end

	dp_0
end


max_profit([1, 2, 3, 0, 2])
