
def heap_sort(arr)
  buildMaxHeap(arr)

  size = arr.size
  i = size - 1
  while i > 0
    arr[0], arr[i] = arr[i], arr[0]
    size -= 1
    i -= 1
    heapify(arr, 0, size)
  end

  arr
end

def buildMaxHeap(arr)
  size = arr.size

  i = size/2
  while i >= 0
    heapify(arr, i, size)
    i-=1
  end
end

def heapify(arr, c, size)
  root = c
  left = 2*c+1
  right = 2*c+2

  if left < size && arr[left] > arr[root]
    root = left
  end

  if right < size && arr[right] > arr[root]
    root = right
  end

  if root != c
    arr[root], arr[c] = arr[c], arr[root]
    heapify(arr, root, size)
  end
end
puts heap_sort([4,3,5,9])