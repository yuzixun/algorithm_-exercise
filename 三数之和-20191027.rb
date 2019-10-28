# @param {Integer[]} nums
# @return {Integer[][]}
def three_sum(nums)
  nums.sort!
  result = []
  size = nums.size

  nums.each_with_index do |num, index|
    if index > 0 && nums[index-1] == nums[index]
      next
    end

    target = -num

    i = index+1
    j = size-1
    while i < j
      cur = nums[i] + nums[j]
      puts "target is #{target}, #{i}, #{j}; #{nums[i]} + #{nums[j]} = #{cur}"
      case
      when cur > target
        j -= 1
      when cur < target
        i += 1
      when cur == target
        result.push([num, nums[i], nums[j]])
        while i < j && nums[i] == nums[i+1]
          i+=1
        end
        while i < j && nums[j] == nums[j-1]
          j -= 1
        end
        i += 1
        j -= 1
      end
    end
  end

  # puts "result is #{result}"
  result
end


nums = [-1,0,1,2,-1,-4]
three_sum(nums)