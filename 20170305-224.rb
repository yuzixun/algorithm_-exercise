# author: yuzixun
# time O(n); space O(n)
# @param {String} s
# @return {Integer}

def calculate(s)
  s.gsub!(' ', '')

  stack = []
  data = 0
  symbol = 1
  result = 0

  s.each_char do |element|
    case element
    when '+'
      result += symbol*data
      symbol = 1
      data = 0
    when '-'
      result += symbol*data
      symbol = -1
      data = 0
    when '('
      stack.push(result, symbol)
      result = 0
      symbol = 1
    when ')'
      previous_result, previous_symbol = stack.pop(2)
      result = previous_result + previous_symbol * (result + symbol*data)
      data = 0
      symbol = 1
    else
      data = data*10+element.to_i
    end
  end
  result += symbol*data

  result
end