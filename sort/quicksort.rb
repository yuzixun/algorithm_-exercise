
# def quick_sort(arr, left, right)

#   if left < right
#     q = partition(arr, left, right)

#     quick_sort(arr, left, q-1)
#     quick_sort(arr, q+1, right)
#   end

#   arr
# end

# def partition(arr, left, right)
#   k = arr[right]
#   i = left-1

#   for j in (left..right) do
#     if arr[j] < k
#       i+=1
#       arr[i], arr[j] = arr[j], arr[i]
#     end
#   end

#   i+=1
#   arr[i], arr[right] = arr[right], arr[i]

#   i
# end

def quick_sort(arr, left, right)
  size = arr.size

  if left < right
    partition_index = partition(arr, left, right)
    quick_sort(arr, left, partition_index)
    quick_sort(arr, partition_index+1, right)
  end

  return arr
end

def partition(arr, left, right)
  q = arr[left]

  index, i = left+1,left+1
  while i <= right
    if q > arr[i]
      arr[i], arr[index] = arr[index], arr[i]
      index+=1
    end

    i+=1
  end

  arr[left], arr[index-1] = arr[index-1], arr[left]
  return index-1
end



puts quick_sort([4,3,5,9], 0, 3)