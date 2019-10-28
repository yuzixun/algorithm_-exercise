
def dfs(s, n, ip, ip_arr)
  if n == 0
    if s == ""
      ip_arr.push(ip[0..-2])
      return
    end
  end

  for i in (1..3) do
    next if i > s.size

    sub_str = s[0, i]
    next if sub_str[0]=="0" && sub_str != '0'

    if sub_str.to_i <= 255
      dfs(s[i..-1] || "", n-1, ip+sub_str+".", ip_arr)
    end
  end
end
# @param {String} s
# @return {String[]}
def restore_ip_addresses(s)
  result = []

  dfs(s, 4, '', result)

  result.uniq
end




def restore_ip_addresses(s)
  result = []
  size = s.size

  for i in (1..[3, size-3].min) do
    for j in (i+1..[i+4, size-2].min) do
      for k in (j+1..[j+4, size-1].min) do
        s1, s2, s3, s4 = s[0...i], s[i...j], s[j...k], s[k...size]
        puts (s1 + "." + s2 + "." + s3 + "." + s4)

        if is_valid?(s1) && is_valid?(s2) && is_valid?(s3) && is_valid?(s4)
          result.push(s1 + "." + s2 + "." + s3 + "." + s4)
        end

      end

    end

  end
  result
end

def is_valid?(s)
  size = s.size
  if size > 3 || size == 0 || (s[0] == '0' && size > 1) || s.to_i > 255
    return false
  end

  true
end


# s = "25525511135"
# s = "0000"
# s = "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
puts "restore_ip_addresses(s) is #{restore_ip_addresses(s)}"
