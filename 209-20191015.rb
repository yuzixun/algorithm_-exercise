# @param {Integer} s
# @param {Integer[]} nums
# @return {Integer}
def min_sub_array_len(s, nums)
  size = nums.size
  min = size +1
  slow = 0
  fast = 0

  sum = 0
  for fast in 0...size do
    sum += nums[fast]

    while sum >= s
      min = [min, fast-slow+1].min
      sum -= nums[slow]
      slow += 1
    end

  end

  min == size + 1 ? 0 : min
end


s, nums = 7, [2,3,1,2,4,3]
# s, nums = 100, []
# s, nums = 11, [1,2,3,4,5]
puts min_sub_array_len(s, nums)
