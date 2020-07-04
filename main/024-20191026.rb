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
def swap_pairs(head)
  pre, pre.next = ListNode.new(0), head
  result = pre
  while (a = pre.next) and (b = a.next)
    pre.next, b.next, a.next = b, a, b.next
    pre = a
  end

  result.next
end