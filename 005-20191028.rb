# @param {String} s
# @return {String}
def longest_palindrome(s)
    size = s.size
    result = ''

    size.times do |i|
      s1 = helper(s, i, i+1)
      s2 = helper(s, i, i)

      if result.size < s1.size
        result = s1
      end
      if result.size < s2.size
        result = s2
      end
    end
    result
end

def helper(s, left, right)
  max_right = s.size
  while left >=0 && max_right > right && s[left] == s[right]
    left -= 1
    right +=1
  end

  return s[left+1..right-1]
end

puts longest_palindrome("cbbd")