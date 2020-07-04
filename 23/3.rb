# @param {Integer} radius
# @param {Integer} x_center
# @param {Integer} y_center
# @param {Integer} x1
# @param {Integer} y1
# @param {Integer} x2
# @param {Integer} y2
# @return {Boolean}


def check_overlap(radius, x_center, y_center, x1, y1, x2, y2)
  rr = radius ** 2
  x = x1
  while x <= x2
    y = y1
    while y <= y2
      if ((x - x_center)**2 + (y - y_center)**2) <= rr
        return true
      end
      y+=1
    end
    x+=1
  end

  (x1 <= x_center && x_center <= x2 && y1-radius <= y_center && y_center <= y2+radius) || (y1 <= y_center && y_center <= y2 && x1-radius <= x_center && x_center <= x2+radius)
end

p check_overlap(1, 0, 0, 1, -1, 3, 1)
p check_overlap(1, 0, 0, -1, 0, 0, 1)
p check_overlap(1, 1, 1, -3, -3, 3, 3)
p check_overlap(1, 1, 1, 1, -3, 2, -1)
p check_overlap(1415, 807, -784, -733, 623, -533, 1005)
p check_overlap(1446, -3296, 4816, 6729, 6217, 9758, 9758)
p check_overlap(34,0,67,32,29,33,31)

