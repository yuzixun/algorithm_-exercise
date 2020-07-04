# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

# @param {ListNode} l1
# @param {ListNode} l2
# @return {ListNode}
def add_two_numbers(l1, l2)
  head = result = ListNode.new(0)

  taken = 0
  while l1 && l2
    cur = l1.val + l2.val + taken
    result.next = ListNode.new(cur%10)
    taken = cur/10
    # puts "taken is #{taken}, cur is #{cur}, l1.val is #{l1.val}， #{l2.val}"
    l1 = l1.next
    l2 = l2.next
    result = result.next
  end

  while taken != 0
    # puts "l1 is #{l1}, l2 is #{l2}"
    case
    when l1
      cur = taken+l1.val
      result.next = ListNode.new(cur%10)
      taken = cur/10
      l1 = l1.next
    when l2
      cur = taken+l2.val
      result.next = ListNode.new(cur%10)
      taken = cur/10
      l2 = l2.next
    else
      result.next = ListNode.new(taken)
      taken = taken/10
    end

    result = result.next
  end

  result.next = l1 || l2
  head.next
end


head1 = pre = ListNode.new(0)
[9].each do |i|
  pre.next = ListNode.new(i)
  pre = pre.next
end

head2 = pre = ListNode.new(0)
[9,9].each do |i|
  pre.next = ListNode.new(i)
  pre = pre.next
end

puts add_two_numbers(head2.next, head1.next).inspect
