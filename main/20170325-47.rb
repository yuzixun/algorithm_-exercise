# @param {Integer[]} nums
# @return {Integer[][]}
#
def permute_unique(nums)
  results = [[]]
  nums.each do |num|
    list = []
    results.each do |result|
      (result.size+1).times do |index|
        list.push(result[0...index]+[num]+result[index..-1])
      end
    end
    results = list
  end
  results.uniq
end


puts "#{permute_unique([1,1,2,1])}"
puts "#{permute_unique([1,2,3])}"
puts "#{permute_unique([3,3,0,3])}"