# Definition for an interval.
class Interval
  attr_accessor :start, :end
  def initialize(s=0, e=0)
    @start = s
    @end = e
  end
end

# @param {Interval[]} intervals
# @return {Integer}
def erase_overlap_intervals(intervals)
  return 0 if intervals.size == 0

  result = 0; _end = nil

  intervals.sort_by!(&:start).each do |interval|
    if _end.nil?
      _end = intervals.first.end
      next
    end

    if interval.start < _end
      _end = [interval.end, _end].min
      result += 1
    else
      _end = interval.end
    end
  end

  result
end


# intervals = [ [1,2], [2,3], [3,4], [1,3] ].map { |s,e| Interval.new(s,e) }
# intervals = [ [1,2], [1,2], [1,2] ].map { |s,e| Interval.new(s,e) }
# intervals = [ [1,2], [2,3] ].map { |s,e| Interval.new(s,e) }
# intervals = [[0,2],[1,3],[2,4],[3,5],[4,6]].map { |s,e| Interval.new(s,e) }
intervals = [[1,2],[2,3],[3,4],[-100,-2],[5,7]].map { |s,e| Interval.new(s,e) }
puts erase_overlap_intervals(intervals)