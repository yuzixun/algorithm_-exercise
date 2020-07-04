# @param {Integer} lo
# @param {Integer} hi
# @param {Integer} k
# @return {Integer}
def get_kth(lo, hi, k)
  return lo if lo == hi
  @cache = {}
  def helper(x)
    @cache[x] ||= begin
      if x == 1
        return 0
      end

      if x%2 == 1
        helper(3 * x + 1)+1
      else
        helper(x / 2)+1
      end
    end
  end

  (lo..hi).to_a.sort do |a, b|
    ha = helper(a)
    hb = helper(b)
    if ha > hb
      1
    elsif ha < hb
      -1
    else
      a > b ? 1  : -1
    end
  end[k-1]
end




lo = 12; hi = 15; k = 2;
puts get_kth(lo, hi, k)

lo = 1; hi = 1; k = 1;
puts get_kth(lo, hi, k)

lo = 7; hi = 11; k = 4;
puts get_kth(lo, hi, k)

lo = 10; hi = 20; k = 5;
puts get_kth(lo, hi, k)

# lo = 569; hi = 571; k = 1;
lo = 1; hi = 1000; k = 777;
puts get_kth(lo, hi, k)
