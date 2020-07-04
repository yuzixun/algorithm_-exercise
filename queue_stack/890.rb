# @param {String} s
# @return {String}
def decode_string(s)
  result = ""
  stack = []
  num =""
  s.each_char do |i|
    if i == "["
      stack.push([num, result])
      num = ""
      result = ""
    elsif i == "]"
      count, last = stack.pop
      result = last+result*count.to_i
    elsif i >= "0" && i<="9"
      num.concat(i)
    else
      result.concat(i)
    end
  end

  result
end

s = "3[a]2[bc]" # , 返回 "aaabcbc".
p decode_string(s)

s = "3[a2[c]]" # , 返回 "accaccacc".
p decode_string(s)

s = "2[abc]3[cd]ef" # , 返回 "abcabccdcdcdef".
p decode_string(s)