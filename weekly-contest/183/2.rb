# @param {String} s
# @return {Integer}
def num_steps(s)
  count = 0
  take = false
  while s != "1"
    if s[-1] == "0"
      if take
        # take = false
        # s[-1] = "0"
        take = false
        # s = s[0..-2]
        # count+=1
        s[-1] = "1"
      else
        take = false
        s = s[0..-2]
        count+=1
      end
    else
      if take
        s = s[0..-2]
        count+=1
      else
        take = true
        # s[-1] = "0"
        s = s[0..-2]
        count+=2
      end
    end
  end
  count+=1 if take
  count
end

p num_steps("1101")
p num_steps("10")
p num_steps("1")