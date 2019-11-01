# @param {Integer[]} nums
# @return {Integer[]}
def next_greater_elements(nums)
	helper = []
	hash = Hash.new(-1)

	(nums+nums).each_with_index do |num, index|
		while !helper.empty? && helper[-1][1] < num
			i, n = helper.pop
			hash[i] = num
		end

		helper.push([index, num])
	end

	nums.size.times.map { |i| hash[i] }
end

nums = [1,2,1]
puts next_greater_elements(nums)