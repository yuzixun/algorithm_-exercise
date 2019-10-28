
def bsearch(nums, target, left, right)
  while left <= right
    mid = left + (right-left)/2
    num = nums[mid]
    if (nums[mid] < nums[0]) == (target < nums[0])
      num = nums[mid]
    else
      num = target < nums[0] ? -1<<32 : 1<<32
    end

    case
    when num < target
      left = mid + 1
    when num > target
      right = mid - 1
    else
      return mid
    end
  end

  -1
end
# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer}
def search(nums, target)
  bsearch(nums, target, 0, nums.size-1)
end

# nums = [4,5,6,7,0,1,2]
# target = 0
# puts search(nums, target)


# nums = [4,5,6,7,0,1,2]
# target = 3
# puts search(nums, target)


# nums = [3,1]
# target = 1
# puts search(nums, target)

nums = [4,5,6,7,0,1,2]
target = 2
puts search(nums, target)