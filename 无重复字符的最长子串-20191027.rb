# @param {String} s
# @return {Integer}
def length_of_longest_substring(s)
  result = current = 0
  hash = {}

  s.each_char.with_index do |c, index|
    if (old_index = hash[c])

      hash.delete_if do |k, i|
        delete = i <= old_index
        current -= 1 if delete
        delete
      end
    end

    hash[c] = index
    current += 1
    result = [result, current].max
  end

  result
end

s = "aab"
s = "pwwkew"
s = "abcabcbb"
s = "dvdf"
puts length_of_longest_substring(s)