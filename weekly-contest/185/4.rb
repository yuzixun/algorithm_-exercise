# @param {Integer} n
# @param {Integer} m
# @param {Integer} k
# @return {Integer}
def num_of_arrays(n, m, k)
  @tmp = []
  (n+1).times do |i|
    ma = []
    (m+1).times do |j|
      ma.push([-1] * (k+1))
    end
    @tmp.push(ma)
  end

  def helper(n,i,k)
    # p "n is #{n}, i is #{i}, k is #{k}, tmp is #{@tmp}"
    if @tmp[n][i][k] != -1
      return @tmp[n][i][k]
    end

    if n == 0  || i == 0 || k ==0
      @tmp[n][i][k] = 0
      return 0
    end

    if n == 1 && k == 1
      @tmp[n][i][k] = 1
      return 1
    end

    r = 0
    for j in (1...i) do
      r+=helper(n-1,j,k-1)
      r%=1000000007
    end

    r+= helper(n-1,i,k) *i
    r%=1000000007
    @tmp[n][i][k] = r
    return r
  end

  r = 0
  for i in (1...m+1) do
    r += helper(n,i,k)
    r %= 1000000007
  end

  return r
end

n = 2
m = 3
k = 1
p num_of_arrays(n, m, k)

n = 5
m = 2
k = 3
p num_of_arrays(n, m, k)

n = 9
m = 1
k = 1
p num_of_arrays(n, m, k)

n = 50
m = 100
k = 25
p num_of_arrays(n, m, k)

n = 37
m = 17
k = 7
p num_of_arrays(n, m, k)