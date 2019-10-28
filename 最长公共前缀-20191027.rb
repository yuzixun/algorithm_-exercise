
# @param {String[]} strs
# @return {String}
def longest_common_prefix(strs)
  return "" if strs.empty?

  # size = strs.map(&:length).min
  # size = strs.select { |str| [min, str.min].min }
  min = -1
  min_str = ""
  strs.each do |str|
    min = [min, str.length].min
    min_str = str
  end

  min_str.each_char.with_index do |c, index|
    strs.each do |_str|
      if _str[index] != c
        return min_str[0,index]
      end
    end
  end

  min_str
end


# strs = ["flower","flow","flight"]
# strs = ["aca","cba"]
strs = ["aa","a"]
# strs = [""]
puts longest_common_prefix(strs)