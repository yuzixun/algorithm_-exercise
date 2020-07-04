# @param {Integer[]} nums
# @return {Integer}
require "set"
def find_maximum_xor(nums)
	result = 0
	mask = 0
	(0..31).reverse_each do |i|
		mask |= (1<<i)

		s = Set.new()
		nums.each do |num|
			s.add(num & mask)
		end
		temp = result | (1<<i)

		s.each do |item|
			if s.include?(temp ^ item)
				result = temp
				break
			end
		end
	end

	result
end

# puts find_maximum_xor([3, 10, 5, 25, 2, 8])
puts find_maximum_xor([4,6,7])