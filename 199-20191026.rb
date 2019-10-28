# Definition for a binary tree node.
# class TreeNode
#     attr_accessor :val, :left, :right
#     def initialize(val)
#         @val = val
#         @left, @right = nil, nil
#     end
# end

def dfs(root, depth, result)
  return if !root

  if result.size == depth
    result.push(root.val)
  end
  dfs(root.right, depth + 1, result)
  dfs(root.left, depth + 1, result)
end

# @param {TreeNode} root
# @return {Integer[]}
def right_side_view(root)
  result = []

  dfs(root, 0, result)

  result
end

