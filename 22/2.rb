# @param {Integer} n
# @param {Integer[][]} reserved_seats
# @return {Integer}
def max_number_of_families(n, reserved_seats)
  count = 0

  result = Hash.new(0)
  reserved_seats.each do |seat|
    row, column = seat

    if column == 2 || column ==3
      result[row] |= 1
    end

    if column == 4 || column == 5
      result[row] |= 3
    end

    if column == 6 || column == 7
      result[row] |= 6
    end

    if column == 8 || column == 9
      result[row] |= 4
    end
  end
  count = (n-result.size) * 2
  result.each do |_, num|
    if num != 7
      count += 1
    end
  end
  count
end