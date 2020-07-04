# # @param {String} num1
# # @param {String} num2
# # @return {String}
# def multiply(num1, num2)
#   res = 0

#   num1.reverse.each_char.with_index do |c1, i|
#     num2.reverse.each_char.with_index do |c2, j|
#       res += (c1.to_i * c2.to_i) * (10 ** (j+i))
#     end
#   end

#   res.to_s
# end


# @param {String} num1
# @param {String} num2
# @return {String}
def multiply(num1, num2)
  k =
  result = ['0'] * (num1.size + num2.size)

  (0..num1.size-1).reverse_each do |i|
    (0..num2.size-1).reverse_each do |j|
      c1 = num1[i]
      c2 = num2[j]
      low_index, high_index = i+j+1, i+j
      tmp = c1.to_i * c2.to_i + result[low_index].to_i

      result[low_index] = (tmp%10).to_s
      result[high_index] = (result[high_index].to_i + (tmp/10)).to_s
    end
  end

  result.each_with_index do |v, i|
    if v != '0'
      return result[i..-1].join
    end
  end

  '0'
end


puts multiply('2', '3')
puts multiply('123', '456')
puts multiply('999', '999')