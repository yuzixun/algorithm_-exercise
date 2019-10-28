# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

def merge(l1, l2)
  dummy = node = ListNode.new(0)

  # puts l1.inspect
  # puts l2.inspect
  while l1 && l2
    if l1.val > l2.val
      node.next, node, l2 = l2, l2, l2.next
    else
      node.next, node, l1 = l1, l1, l1.next
    end
  end

  node.next = l1 || l2

  dummy.next
end
# @param {ListNode} head
# @return {ListNode}
def sort_list(head)
  if head.nil? || head.next.nil?
    return head
  end

  pre, slow, fast = nil, head, head
  while fast && fast.next
    pre, slow, fast = slow, slow.next, fast.next.next
  end
  pre.next = nil

  merge(sort_list(head), sort_list(slow))
end

head = pre = ListNode.new(0)
[4,2,1,3].each do |i|
  pre.next = ListNode.new(i)
  # puts pre.inspect
  pre = pre.next
end
puts  sort_list(head.next).inspect

# def quick_sort(arr, i, j)
#   if i < j
#     q = partition(arr, i, j)

#     quick_sort(arr, i, q-1)
#     quick_sort(arr, q+1, j)
#   end

#   arr
# end

# def partition(arr, left, right)
#   x = arr[right]
#   i = left - 1

#   for j in (left..right-1) do

#     if arr[j] < x
#       i+=1
#       arr[i], arr[j] = arr[j], arr[i]
#     end
#     puts "arr is #{arr}, i is #{i}, j is #{j}"
#   end

#   i+=1
#   arr[i], arr[right] = [arr[right], arr[i]]
#   i
# end


# puts quick_sort([6, -5, 5, 3, 4, 0, -1], 0, 4)