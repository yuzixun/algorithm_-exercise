# @param {Integer[][]} rooms
# @return {Boolean}
require 'set'
def can_visit_all_rooms(rooms)
	visited = Set.new([0])
	queue = rooms[0]

	while !queue.empty?
		cur = queue.shift
		# puts "queue is #{queue}, cur is #{cur}"
		next if queue.include?(cur)
		visited.add(cur)
		queue.concat(rooms[cur])
		rooms[cur] = []
	end

	# puts "queue is #{rooms}, visited is #{visited}"
	rooms.size == visited.size
end

puts can_visit_all_rooms([[1],[2],[3],[]])
puts can_visit_all_rooms([[1,3],[3,0,1],[2],[0]])
puts can_visit_all_rooms([[1],[1,1]])