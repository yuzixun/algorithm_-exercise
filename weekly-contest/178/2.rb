# @param {String[]} votes
# @return {String}
def rank_teams(votes)
	if votes.size == 1
		return votes[0]
	end

	# 名词 -> 单词 -> 个数
	hash = {}
	votes.each_with_index do |vote, vote_index|
		vote.each_char.with_index do |word, word_index|
			hash[word_index] ||= {}
			hash[word_index][word] ||= 0
			hash[word_index][word] += 1
		end
	end

	result = ""
	max_arr = []
	hash.size.times do |key|
		max = 0
		hash[key].each do |word, value|
			if max < value
				max_arr[key] = [word]
			elsif max == value
				max_arr[key].push(word)
			end
		end
		# if max_arr.size == 1
		# 	result.concat(max_arr[0])
		# end
	end

	puts "max_arr is #{max_arr}"
end

puts rank_teams(["ABC", "ACB", "ABC", "ACB", "ACB"])
puts rank_teams(["WXYZ", "XYZW"])
puts rank_teams(["ZMNAGUEDSJYLBOPHRQICWFXTVK"])
puts rank_teams(["BCA", "CAB", "CBA", "ABC", "ACB", "BAC"])
puts rank_teams(["M", "M", "M", "M"])