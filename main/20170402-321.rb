# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @param {Integer} k
# @return {Integer[]}

def get_max_from(nums, k)
  nums_size = nums.size
  result = []

  nums_size.times do |index|
    num = nums[index]

    while (!result.empty?) && result.size + nums_size - index > k && result.last < num
      result.pop
    end

    result.push(num) if result.size < k
  end

  result
end

def merge(nums1, nums2)
  result = []

  while (!nums1.empty?) || (!nums2.empty?)
    result.push([nums1, nums2].max == nums1 ? nums1.shift : nums2.shift)
  end

  result
end

def max_number(nums1, nums2, k)
  begin
    ([0, k-nums2.size].max..[k, nums1.size].min).map do |i|
      merge(get_max_from(nums1, i), get_max_from(nums2, k-i))
    end
  end.max
end


# puts "#{max_number([3, 4, 6, 5], [9, 1, 2, 5, 8, 3], 5)}"
# puts "#{max_number([6, 7], [6, 0, 4], 5)}"
# puts "#{max_number([3, 9], [8, 9], 3)}"
# puts "#{max_number([9,1,2,5,8,3], [3,4,6,5], 5)}"

a = [6,3,9,0,5,6]
b = [2,2,5,2,1,4,4,5,7,8,9,3,1,6,9,7,0]
k = 23
puts "#{max_number(a, b, k)}"