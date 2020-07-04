# @param {Integer} n
# @param {Integer} head_id
# @param {Integer[]} manager
# @param {Integer[]} inform_time
# @return {Integer}
def num_of_minutes(n, head_id, manager, inform_time)
  cache = {}

  get_value = Proc.new do |index|
    result = cache[index] || begin
      (index == -1 || head_id == index) ? inform_time[head_id] : (inform_time[index] + get_value.call(manager[index]))
      # puts " index is #{index}, inform_time[index] is #{inform_time[index]}, r is #{r}"
    end
    # puts "cache si #{cache}, result is #{result}, index is #{index}"
    cache[index] ||= result
    result
  end

  max = -1
  n.times do |index|
    manager[index] = get_value.call(index)
    if max < manager[index]
      max = manager[index]
    end
  end
  # puts "hash is #{cache}, max is #{max}"
  max
end




# n = 1
# head_id = 0
# manager = [-1]
# inform_time = [0]
# puts num_of_minutes(n, head_id, manager, inform_time)

n = 6
head_id = 2
 manager = [2,2,-1,2,2,2]
 inform_time = [0,0,1,0,0,0]
puts num_of_minutes(n, head_id, manager, inform_time)

n = 7
head_id = 6
 manager = [1,2,3,4,5,6,-1]
 inform_time = [0,6,5,4,3,2,1]
puts num_of_minutes(n, head_id, manager, inform_time)

n = 15
head_id = 0
 manager = [-1,0,0,1,1,2,2,3,3,4,4,5,5,6,6]
 inform_time = [1,1,1,1,1,1,1,0,0,0,0,0,0,0,0]
puts num_of_minutes(n, head_id, manager, inform_time)

n = 4
head_id = 2
 manager = [3,3,-1,2]
 inform_time = [0,0,162,914]
puts num_of_minutes(n, head_id, manager, inform_time)
