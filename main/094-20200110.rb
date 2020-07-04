def inorder_traversal(root)
  result = []
  white, black = 0, 1
  stack = [[white, root]]

  while !stack.empty?
    color, node = stack.pop

    next if !node

    if color == white
      stack.push([white, node.right]) # if node.right
      stack.push([black, node]) # if node
      stack.push([white, node.left]) # if node.left
    else
      result.push(node.val)
    end
    # puts stack.inspect
  end

  result
end
