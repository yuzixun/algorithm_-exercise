
def quick_sort(arr, left, right)

  if left < right
    q = partition(arr, left, right)

    quick_sort(arr, left, q-1)
    quick_sort(arr, q+1, right)
  end

  arr
end

def partition(arr, left, right)
  k = arr[right]
  i = left-1

  for j in (left..right) do
    if arr[j] < k
      i+=1
      arr[i], arr[j] = arr[j], arr[i]
    end
  end

  i+=1
  arr[i], arr[right] = arr[right], arr[i]

  i
end

puts quick_sort([4,3,5,9], 0, 3)