# @param {String} s
# @param {String[]} d
# @return {String}
def find_longest_word(s, d)
  d.sort! { |a, b| a.size == b.size ? a <=> b : b.size <=> a.size }

  i = j = 0

  s_size = s.size
  d.each do |w|
    w_size = w.size
    while j != w_size && i != s_size
      if s[i] == w[j]
        i += 1; j += 1
      else
        i += 1
      end
    end

    return w if j == w.size
    i = j = 0
  end

  ''
end

s = "abpcplea"
d = ["ale","apple","monkey","plea"]
puts find_longest_word(s, d)