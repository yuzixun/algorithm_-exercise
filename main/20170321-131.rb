# @param {String} s
# @return {String[][]}

def is_palindrome(string, start_pos, end_pos)
  temp = string[start_pos..end_pos]
  temp == temp.reverse
end

def recur(s, start_pos, result, results)
  if start_pos == s.size
    results.push(result.dup)
    result.clear
  end

  for index in (start_pos..s.size-1)
    if is_palindrome(s, start_pos, index)
      result.push(s[start_pos..index])
      recur(s, index+1, result, results)
    end
  end

end

def partition(s)
  result = []
  results = []

  recur(s, 0, result, results)
  results
end


a = 'aab'

puts "#{partition(a)}"