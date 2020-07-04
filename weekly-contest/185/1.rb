# @param {String} s
# @return {String}
def reformat(s)
  nums = []
  chars = []
  s.each_char do |c|
    if c >= 'a' && c<='z'
      chars.push(c)
    else
      nums.push(c)
    end
  end
  # numIndex = charIndex = 0
  # while numIndex < s.size && charIndex < s.size
  #   while s[charIndex]  >= 'a' && s[charIndex] <= 'z'
  #     charIndex+=1
  #   end

  #   while s[numIndex]  >= '0' && s[numIndex] <= '9'
  #     numIndex+=1
  #   end
  # end

  return "" if (chars.size - nums.size).abs > 1
  char_uniq = chars.uniq.size
  return "" if char_uniq == 1 && char_uniq == nums.uniq.size

  arr = []
  i = j = 0

  while i < chars.size && j < nums.size
    if i == j
      if chars.size > nums.size
        arr.push(chars[i])
        i+=1
      else
        arr.push(nums[j])
        j+=1
      end

    elsif i < j
      arr.push(chars[i])
      i+=1
    else
      arr.push(nums[j])
      j+=1
    end
    # puts "i is #{i}, j is #{j}, chars is #{chars}, nums is #{nums}, #{arr}"
  end
  # puts "#{i} , #{chars.size}, #{j}, #{nums.size}"
  if i <= chars.size-1
    arr.push(chars[i])
    i+=1
  end
  if j <= nums.size-1
    arr.push(nums[j])
    j+=1
  end
  result = arr.join
  result == s ? result.reverse : result
end

p reformat("a0b1c2")
p reformat("leetcode")
p reformat("1229857369")
p reformat("covid2019")
p reformat("ab123")