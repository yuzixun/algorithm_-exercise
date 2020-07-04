# @param {Integer[]} nums
# @return {String}
def largest_number(nums)
  nums.map(&:to_s).sort { |a, b| "#{b}#{a}" <=> "#{a}#{b}" }.join.to_i.to_s
end

puts largest_number([3, 30, 34, 5, 9])