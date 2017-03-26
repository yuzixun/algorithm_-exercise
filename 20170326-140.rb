# @param {String} s
# @param {String[]} word_dict
# @return {String[]}
def dfs(string, word_dict, map)
  return map[string] if map.member?(string)

  res = []

  if string.empty?
    res.push('')
    return res
  end

  word_dict.each do |word|
    if string.start_with?(word)
      sublist = dfs(string[word.size..-1], word_dict, map)
      sublist.each do |sub|
        res.push (word + ' ' + sub).strip
      end
    end
  end

  map[string] = res
  return res
end

def word_break(s, word_dict)
  dfs(s, word_dict, {})
end



s = 'catsanddog'
dict = ["cat", "cats", "and", "sand", "dog"]
puts "all result is #{word_break(s, dict)}"
s = "aaaaaaa"
dict = ["aaaa", "aaa"]
puts "all result is #{word_break(s, dict)}"
s = "bb"
dict = ["a","b","bbb","bbbb"]
puts "all result is #{word_break(s, dict)}"