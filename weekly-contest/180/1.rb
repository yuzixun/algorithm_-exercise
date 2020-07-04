# @param {Integer[][]} matrix
# @return {Integer[]}
def lucky_numbers (matrix)
  return 0 if matrix.empty?

  result = []
  matrix.each do |row|
    min = row.min
    row.find_all.with_index do |item, index|
      if item == min
        max = matrix.map { |columns| columns[index] }.max
        # puts " min is #{min}, max is #{max}"
        result.push(min) if min == max
      end
    end
  end

  result.empty? ? nil : result
end

puts lucky_numbers([[3,7,8],[9,11,13],[15,15,17]])
puts lucky_numbers([[1,10,4,2],[9,3,8,7],[15,16,17,12]])
puts lucky_numbers([[7,8],[1,2]])
puts lucky_numbers([[36376,85652,21002,4510],[68246,64237,42962,9974],[32768,97721,47338,5841],[55103,18179,79062,46542]])