# @param {Integer[][]} points
# @return {Integer}
def find_min_arrow_shots(points)
  count = 0

  points.sort! { |a, b| a.last <=> b.last }

  last = nil
  points.each do |point|
    next if (!last.nil?) && last >= point.first

    last = point.last
    count += 1
  end

  count
end



puts find_min_arrow_shots([[10,16], [2,8], [1,6], [7,12]])
puts find_min_arrow_shots([[9,12],[1,10],[4,11],[8,12],[3,9],[6,9],[6,7]])

