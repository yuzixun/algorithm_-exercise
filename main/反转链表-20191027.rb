# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

# @param {ListNode} head
# @return {ListNode}
def reverse_list(head)
  pre, cur = nil, head

  while cur
    cur.next, pre, cur = pre, cur, cur.next
  end

  pre
end

head = pre = ListNode.new(0)
[4,2,1,3].each do |i|
  pre.next = ListNode.new(i)
  pre = pre.next
end
puts reverse_list(head.next).inspect
