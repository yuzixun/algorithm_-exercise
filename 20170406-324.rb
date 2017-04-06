# @param {Integer[]} nums
# @return {Void} Do not return anything, modify nums in-place instead.
def wiggle_sort(nums)
  return if nums.size <= 1
  nums.sort!
  half = nums.size/2

  if nums.size.even?
    smaller, larger = nums[0..half-1].reverse, nums[half..-1].reverse
  else
    smaller, larger = nums[0..half].reverse, nums[half+1..-1].reverse
  end

  smaller.size.times do |index|
    c_index = 2*index

    unless smaller[index].nil?
      nums[c_index] = smaller[index]
      c_index += 1
    end

    nums[c_index] = larger[index] unless larger[index].nil?
  end
end


wiggle_sort([1, 5, 1, 1, 6, 4])
wiggle_sort([1, 3, 2, 2, 3, 1])
wiggle_sort([1, 8, 5, 6, 0, 7, 9, 2, 4, 3])
wiggle_sort([5, 3, 0, 2, 1, 8, 4, 7, 6])
wiggle_sort([1,1,2,1,2,2,1])
wiggle_sort([5,3,1,2,6,7,8,5,5])
wiggle_sort([1])