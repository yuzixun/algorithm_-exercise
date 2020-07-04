# @param {Integer[]} nums
# @return {Integer[]}
def min_subsequence(nums)
  nums.sort! { |a,b| b<=>a}
  all_sum = nums.reduce(:+)
  result = []
  left = 0
  right = all_sum
  nums.each do |num|
    break if left > right
    left += num
    right -= num
    result.push(num)
  end

  result
end

p min_subsequence([4,4,7,6,7])
p min_subsequence([6])
p min_subsequence([])
p min_subsequence([4,3,10,9,8])