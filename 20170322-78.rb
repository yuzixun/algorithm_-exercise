# @param {Integer[]} nums
# @return {Integer[][]}
def subsets(nums)
  nums.reduce([[]]) { |results, item| results += results.map { |r| r += [item] } }
end

puts "#{subsets([1,2,3])}"