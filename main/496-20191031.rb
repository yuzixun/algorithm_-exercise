# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Integer[]}
def next_greater_element(nums1, nums2)
	l1 = nums1.size
	l2 = nums2.size
	hash = Hash.new(-1)
	helper = []

	nums2.each_with_index do |i2, index|
		while !helper.empty? && helper[-1] < i2
			hash[helper.pop] = i2
		end
		helper.push(i2)
	end

	# result = []
	nums1.map { |i| hash[i] }
end

# nums1 = [4,1,2]
# nums2 = [1,3,4,2]

nums1 = [2,4]
nums2 = [1,2,3,4]
puts next_greater_element(nums1, nums2)