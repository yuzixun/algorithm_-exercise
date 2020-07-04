# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

# @param {ListNode} head
# @return {ListNode}
def reverse_list(head)
  prev, cur = nil, head

  while cur
    cur.next, prev, cur = prev, cur, cur.next
  end

  return cur
end