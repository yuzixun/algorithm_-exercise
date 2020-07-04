
# @param {Integer} n
# @return {Integer}
def fib(n)
  a, b = 0, 1
  n.times { a, b = a + b, a }
  a
end

puts fib(2)
puts fib(3)
puts fib(4)
