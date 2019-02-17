# Given an array of integers, return indices of the two numbers such that they add up to a specific target.
# You may assume that each input would have exactly one solution, and you may not use the same element twice.
# Example:
# Given nums = [2, 7, 11, 15], target = 9,
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].
#


# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
  results = []
  nums.each_with_object([]) do |object, index|
    surplus = target - object
    if (nums - [object]).to_a.include?(surplus)
      results.push(index)
      results.push(nums.index(surplus))
    end
  end
  results
end


# nums = [2, 7, 11, 15]
# target = 9
nums = [3,2,4]
target = 6
puts two_sum(nums, target)