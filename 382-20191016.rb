# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

class Solution

=begin
    @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node.
    :type head: ListNode
=end
    def initialize(head)
      @head = head
      # @rand = Rand.random()
    end


=begin
    Returns a random node's value.
    :rtype: Integer
=end
    def get_random()
      result = @head
      node = @head
      i = 1

      while node
        if Random.new.rand(i) == 0
          result = node
        end
        node = node.next
        i+=1
      end

      result
    end


end

# Your Solution object will be instantiated and called as such:
# obj = Solution.new(head)
# param_1 = obj.get_random()