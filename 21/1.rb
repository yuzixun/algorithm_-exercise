# @param {String} s
# @return {String}
def sort_string(s)

	result = ""
	finding_min = true
	arr = s.each_char.to_a.sort


	while !arr.empty? # .length != used.length
		used = []
		prev = nil
		arr.each_with_index do |l, index|
			next if prev == l || used.include?(index)
			prev = l
			result.concat(l)
			used.push(index)
		end

		# puts "arr is #{arr}, used is #{used}"
		used.reverse.each do |del|
    	arr.delete_at(del)
		end

		finding_min = !finding_min
		arr = arr.reverse
	end
	# puts result
	result
end


# sort_string("spo")
# sort_string("")
# sort_string("rat")
# sort_string("ggggggg")
# sort_string("leetcode")
# sort_string("aaaabbbbcccc")
sort_string("dvhpp")