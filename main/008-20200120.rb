
# @param {String} str
# @return {Integer}
def my_atoi(str)
	# str.lstrip!
	# result = /^[\+\-]?\d+/.match(str).to_s.to_i
	# [[result, 2**31].min, -2**31].max
	[[str.to_i, limit = -2**31].max, -limit - 1].min
end

my_atoi("42")
my_atoi("   -42")
my_atoi("4193 with words")
my_atoi("words and 987")
my_atoi("-91283472332")