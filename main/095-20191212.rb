# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end


def generate_from_to(left, right)
	result = []
	if left > right
		return [nil]
	end

	for i in left..right do
		left_nodes = generate_from_to(left, i-1)
		right_nodes = generate_from_to(i+1, right)
		left_nodes.each do |left_node|
			right_nodes.each do |right_node|
				t = TreeNode.new(i)
				t.left = left_node
				t.right = right_node
				result.push(t)
			end
		end
	end

	result
end

# @param {Integer} n
# @return {TreeNode[]}
def generate_trees(n)
	return [] if n == 0
	generate_from_to(1, n)
end


# puts generate_trees(3)
puts generate_trees(0).inspect