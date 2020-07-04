def codec(root)
  arr = [root]
  helperCode(arr, index)

  wuk
  # result = ""
  arr.map { |s| s&.val }
  # arr.map(&:val)
end

def helperCode(arr, index)
  return if index >= arr.size

  root = arr[index]
  if root
    arr.push(root.left)
    arr.push(root.right)
  end

  helperCode(arr, index+1)
end
