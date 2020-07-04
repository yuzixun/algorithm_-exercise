# @param {Integer[]} start1
# @param {Integer[]} end1
# @param {Integer[]} start2
# @param {Integer[]} end2
# @return {Float[]}
def intersection(start1, end1, start2, end2)
  x, y = 0, 0
  x1, y1 = start1
  x2, y2 = end1
  x3, y3 = start2
  x4, y4 = end2

  delta = Proc.new { |a,b,c,d| a*d-b*c}
  array_sort = Proc.new do |arr1, arr2|
    result = arr1 <=> arr2
    if result == -1
      [arr2, arr1]
    else
      [arr1, arr2]
    end
  end
  d = delta.call(x1-x2, x4-x3, y1-y2, y4-y3)
  p = delta.call(x4-x2, x4-x3, y4-y2, y4-y3)
  q = delta.call(x1-x2, x4-x2, y1-y2, y4-y2)

  if d != 0
    lam, eta = p/d, q/d
    if !(0 <= lam && lam <= 1 && 0 <= lam && lam <= 1)
      return []
    else
      return [(lam * x1 + (1-lam) * x2).to-f, (lam *y1 + (1-lam) * y2).to_f]
    end
  end

  if p != 0 && q!=0
    return []
  end

  t1, t2 = array_sort.call(start1, end1), array_sort.call(start2, end2)
  if (t1[1] <=> t2[0] == -1) || (t2[1] <=> t1[0] == -1)
    return []
  end

  return [t1[0].to_f, t2[0].to_f].max
end