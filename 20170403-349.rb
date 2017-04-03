# @param {Integer[]} nums1
# @param {Integer[]} nums2
# @return {Integer[]}
def intersection(nums1, nums2)
  nums1 & nums2
end

puts intersection([1, 2, 2, 1], [2, 2])