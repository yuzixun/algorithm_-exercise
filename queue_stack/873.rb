# @param {String[]} deadends
# @param {String} target
# @return {Integer}
require 'set'
def open_lock(deadends, target)
	deadends = Set.new(deadends)
	return -1 if deadends.include?("0000")
	set = []
	set.push(["0000", 0])

	while !set.empty?
		str, count = set.shift

		for i in (0..3) do
			[1, -1].each do |delta|
				n_str = str.dup
				n_str[i] = ((str[i].to_i+delta)%10).to_s
				return count+1 if target == n_str
				next if deadends.include?(n_str)
				deadends.add(n_str)
				set.push([n_str, count+1])
			end
		end
	end
	-1
end

# deadends = ["0201","0101","0102","1212","2002"]
# target = "0202"
# p open_lock(deadends, target)
# deadends = ["8888"]
# target = "0009"
# p open_lock(deadends, target)
# deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"]
# target = "8888"
# p open_lock(deadends, target)
# deadends = ["0000"]
# target = "8888"
# p open_lock(deadends, target)