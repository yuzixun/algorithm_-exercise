# @param {String} s
# @param {Integer} length
# @return {String}
def replace_spaces(s, length)
	s[0..length-1].gsub(" ", "%20")
end

puts replace_spaces("Mr John Smith    ", 13)
puts replace_spaces("               ", 5)