# @param {Integer[][]} intervals
# @return {Integer[][]}
def merge(intervals)
  return [] if intervals.empty?
  intervals.sort_by!(&:first)

  results = [intervals[0].clone]
  size = intervals.length

  for i in 1..size-1 do
    prev_min, prev_max = results[-1]
    cur_min, cur_max = intervals[i]

    new_min = new_max = nil
    if prev_min <= cur_min && cur_min <= prev_max
      new_min = [prev_min, cur_min].min
      new_max = [prev_max, cur_max].max
      results[-1] = [new_min, new_max]
    else
      results.push(intervals[i])
    end
  end

  results
end

# intervals = [[1,3],[2,6],[8,10],[15,18]]
# intervals = [[1,4],[4,5]]
intervals = []
merge(intervals)