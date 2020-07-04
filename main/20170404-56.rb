# Definition for an interval.
class Interval
  attr_accessor :start, :end
  def initialize(s=0, e=0)
    @start = s
    @end = e
  end
end

# @param {Interval[]} intervals
# @return {Interval[]}
def merge(intervals)
  result = []; _end = nil

  intervals.sort_by!(&:start).each do |interval|
    if (!_end.nil?) && interval.start <= _end
      if interval.end > _end
        result[-1][-1] = interval.end
        _end = interval.end
      end
    else
      result.push([interval.start, interval.end])
      _end = interval.end
    end
  end

  result
end


array = []
[[1,4],[4,5]].each do |e|
  array.push(Interval.new(e.first, e.last))
end

p merge(array)

array = []
[[1,4],[2,3]].each do |e|
  array.push(Interval.new(e.first, e.last))
end

p merge(array)

array = []
[[1,4],[0,2],[3,5]].each do |e|
  array.push(Interval.new(e.first, e.last))
end

p merge(array)

array = []
[[2,3],[4,5],[6,7],[8,9],[1,10]].each do |e|
  array.push(Interval.new(e.first, e.last))
end

p merge(array)
