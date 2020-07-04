# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

# @param {TreeNode} root
# @return {Integer}
def max_depth(root)
  return 0 if !root
  [max_depth(root.left), max_depth(root.right)].max + 1
end

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

  root
end

root = [3,5,1,6,2,0,8,nil,nil,7,4]
# root = [3,9,20,nil,nil,15,7]
puts max_depth(build_tree(root))