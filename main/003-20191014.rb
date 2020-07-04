# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  max = 0
  cur = 0
  hash = {}

  s.each_char.with_index do |c, i|
    if hash[c] && hash[c] < i
      hash.delete_if { |_, index| index < hash[c] }
      hash[c] = i
      cur = hash.length
    else
      hash[c] = i
      cur += 1
    end

    max = cur if cur > max
  end

  max
end


puts length_of_longest_substring("abcabcbb")
puts length_of_longest_substring("bbbbb")
puts length_of_longest_substring("pwwkew")