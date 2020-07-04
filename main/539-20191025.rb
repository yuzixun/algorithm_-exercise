# @param {String[]} time_points
# @return {Integer}
def find_min_difference(time_points)
  times = []
  time_points.each do |time_point|
    h, m = time_point.split(":")
    times.push(h.to_i*60+m.to_i)
  end

  times.sort!

  min = 1440
  end_index = times.size-1

  for i in (1..end_index) do
    min = [min, times[i]-times[i-1]].min

    if i == end_index
      min = [min, 1440+times[0]-times[i]].min
    end
  end

  min
end
