# @param {String[]} words
# @return {String[]}
def string_matching(words)
  result = []
  words.each_with_index do |word, i|
    result.concat( words.find_all.with_index do |w, fi|
      next if i == fi
      word.include?(w)
    end)
  end

  result.uniq
end

string_matching(["mass","as","hero","superhero"])
string_matching(["leetcode","et","code"])
string_matching(["blue","green","bu"])
string_matching(["leetcoder","leetcode","od","hamlet","am"])