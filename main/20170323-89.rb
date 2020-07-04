# @param {Integer} n
# @return {Integer[]}
def gray_code(n)
  # results = [0]
  # n.times do |index|
  #   puts '-------'
  #   results += results.reverse.map do |e|
  #     a = e + 2**index
  #     puts "#{e.to_s(2)} => #{a.to_s(2)}"
  #     a
  #   end
  # end
  # results
  n.times.reduce([0]) { |results, index| results += results.reverse.map { |e| e + 2**index } }
end

puts "#{gray_code(3)}"