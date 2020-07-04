# @param {Integer[]} nums
# @return {Integer}
def first_missing_positive(nums)
  return 1 if nums.empty?
  h = nums.each_with_object({}) { |n, h| h[n] = true }
  max = h.keys.max
  (1..max).each do |n|
    unless h[n]
      return n
    end
  end
  [max+1, 1].max
end

puts first_missing_positive([3,4,-1,1])
puts first_missing_positive([3,4,-1,1])
puts first_missing_positive([1,2,0])
puts first_missing_positive([7,8,9,11,12])
puts first_missing_positive([0])
puts first_missing_positive([-5])