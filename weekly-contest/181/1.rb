
# @param {Integer[]} nums
# @param {Integer[]} index
# @return {Integer[]}
def create_target_array(nums, index)
  target = []

  index.size.times do |i|
    target.insert(index[i], nums[i])
  end

  target
end


nums = [0,1,2,3,4]; index = [0,1,2,2,1]
p create_target_array(nums, index)
nums = [1,2,3,4,0]; index = [0,1,2,3,0]
p create_target_array(nums, index)