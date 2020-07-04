# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

# def find_path(root, value, path)
#   return false if !root

#   if value == root.val
#     found = true
#   end

#   path.push(root.val)
#   result = found || find_path(root.left, value, path) || find_path(root.right, value, path)

#   if !result
#     path.delete(root.val)
#   end

#   result
# end
# # @param {TreeNode} root
# # @param {TreeNode} p
# # @param {TreeNode} q
# # @return {TreeNode}
# def lowest_common_ancestor(root, p, q)
#   p_path = []
#   q_path = []
#   find_path(root, p, p_path)
#   find_path(root, q, q_path)

#   same = nil
#   min = [p_path.size, q_path.size].min

#   puts "p_path is #{p_path}, q_path is #{q_path}"

#   min.times do |i|
#     if p_path[i] == q_path[i]
#       same = p_path[i]
#     end
#   end

#   same
# end



# @param {TreeNode} root
# @param {TreeNode} p
# @param {TreeNode} q
# @return {TreeNode}
def lowest_common_ancestor(root, p, q)
  return nil if !root

  left = lowest_common_ancestor(root.left, p, q)
  right = lowest_common_ancestor(root.right, p, q)

  # puts "left is #{left&.val}, right is #{right&.val}, val si #{root.val}, q is #{q}, p is #{p}"
  if left && right || root.val == p.val || root.val == q.val
    return root
  else
    return left || right
  end
end

root = [3,5,1,6,2,0,8,nil,nil,7,4]
p = 5
q = 1

def build_tree(arr)
  root = TreeNode.new(arr.shift)

  node_arr = [root]

  while !arr.empty?
    node = node_arr.shift

    left = arr.shift
    if left
      node.left = TreeNode.new(left)
      node_arr.push(node.left)
    end

    right = arr.shift
    if right
      node.right = TreeNode.new(right)
      node_arr.push(node.right)
    end

  end

  # puts "#{root.inspect}"

  root
end

puts lowest_common_ancestor(build_tree(root), p, q).val