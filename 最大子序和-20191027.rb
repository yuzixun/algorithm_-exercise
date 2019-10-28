# @param {Integer[]} nums
# @return {Integer}
def max_sub_array(nums)
  sums = []
  result = nums[0]

  nums.each_with_index do |num, index|
    sums[index] = [sums[index-1].to_i+num, num].max
    if sums[index] > result
      result = sums[index]
    end
  end

  result.to_i
end

puts max_sub_array([-2,1,-3,4,-1,2,1,-5,4])