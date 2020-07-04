# @param {Integer[][]} intervals
# @param {Integer[]} new_interval
# @return {Integer[][]}
def insert(intervals, new_interval)
  new_min, new_max = new_interval
  left, right = [], []

  intervals.each do |interval|
    cur_min, cur_max = interval
    case
    when cur_max < new_min
      left.push(interval)
    when cur_min > new_max
      right.push(interval)
    else
      new_min = [new_min, cur_min].min
      new_max = [new_max, cur_max].max
    end
  end

  return left + [[new_min, new_max]] + right
end

intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]
new_interval = [4,8]
#
# intervals = [[1,3],[6,9]]
# new_interval = [2,5]
# intervals = [[1,5]]
# new_interval = [6,7]

result = insert(intervals, new_interval)
puts "result is #{result}"