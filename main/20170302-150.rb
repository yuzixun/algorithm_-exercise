# @param {String[]} tokens
# @return {Integer}
def eval_rpn(tokens)
  symbols = %w(+ - * /)
  result_array = []

  tokens.each do |value|
    if symbols.include?(value)
      a, b = result_array.pop(2)
      value = a.to_f.send(value, b)
    end

    result_array.push(value.to_i)
  end
  result_array[0]
end