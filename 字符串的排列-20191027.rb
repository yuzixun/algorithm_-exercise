# @param {String} s1
# @param {String} s2
# @return {Boolean}
def check_inclusion(s1, s2)

  hash = Hash.new(0)
  s1.each_char do |c|
    hash[c] += 1
  end

  result = slow = 0
  s2.each_char.with_index do |c, fast|
    hash[c] -= 1

    if hash[c] < 0
      while hash[c] != 0
        hash[s2[slow]] += 1
        slow += 1
      end
    end

    result = [result, fast - slow+1].max
  end

  result == s1.size
end


s1,  s2 = "aboi", "eidbaooo"
puts check_inclusion(s1, s2)
s1,  s2 = "ab", "eidboaoo"
puts check_inclusion(s1, s2)

