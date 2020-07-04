# @param {Integer} n
# @return {String}
def generate_the_string(n)
	return "a" if n == 1

	result = ""
	if n%2 == 1
		result = "a"*(n-2)
		result.concat("b")
		result.concat("c")
	else
		result = "a"*(n-1)
		result.concat("b")
	end

	result
end

puts generate_the_string 1
puts generate_the_string 4
puts generate_the_string 2
puts generate_the_string 3
puts generate_the_string 7
