# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

# @param {TreeNode} root
# @return {Integer[]}
# require "set"
# def inorder_traversal(root)
# 	visited = Set.new([])
# 	stack = [root]
# 	result = []
# 	while !root.empty?
# 		cur = stack.pop
# 		if visited.include?(cur)
# 			result.push(cur)
# 			next
# 		end
# 		visited.add(cur)
# 		stack.push(root.right) if root.right
# 		stack.push(root)
# 		stack.push(root.left) if root.left
# 	end

# 	result
# end

def inorder_traversal(root)
	result = []
	stack = []

	while root || !stack.empty?
		while root
			stack.add(root)
			root = root.left
		end

		if root = stack.pop
			result.push(root.val)
			root = root.right
		end
	end

	return result
end