# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

# @param {ListNode} head
# @param {Integer} k
# @return {ListNode}
def reverse_k_group(head, k)
  new_head = jump = ListNode.new(0)
  new_head.next = l = r = head

  while true
    i = 0

    while i < k && r
      i+=1
      r=r.next
    end

    if i == k
      pre, cur = r, l
      puts "pre, cur is #{pre&.val}, #{cur&.val}"
      k.times {
        cur.next, pre, cur = pre, cur, cur.next
      }

      # puts "pre, cur, l, r is #{pre&.val}, #{cur&.val}, #{l&.val}, #{r&.val}"
      jump.next, jump, l = pre, l, r
    else
      return new_head.next
    end
  end
end


arr = [1,2,3,4,5,6,7,8,9,10,11,12,13]
k = 3

head = cur = ListNode.new(0)
arr.each do |a|
  cur.next = ListNode.new(a)
  cur = cur.next
end

head = reverse_k_group(head.next, k)
while head
  puts head.val
  head = head.next
end

