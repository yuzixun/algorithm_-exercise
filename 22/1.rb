# @param {Integer[]} arr1
# @param {Integer[]} arr2
# @param {Integer} d
# @return {Integer}
def find_the_distance_value(arr1, arr2, d)
  count = 0

  arr1.each do |i|
    if arr2.select { |j| (i-j).abs <= d }.empty?
      count +=1
    end
  end

  count
end

find_the_distance_value([4,5,8], [10,9,1,8],2)
find_the_distance_value([1,4,2,3], [-4,-3,6,10,20,30],3)
find_the_distance_value([2,1,100,3], [-5,-2,10,-3,7],6)