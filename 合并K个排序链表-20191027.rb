# Definition for singly-linked list.
class ListNode
    attr_accessor :val, :next
    def initialize(val)
        @val = val
        @next = nil
    end
end

# @param {ListNode[]} lists
# @return {ListNode}
def merge_k_lists(lists)
  result = head = ListNode.new(0)
  hash = {}
  lists.each do |list|
    next if list.nil?

    hash[list.val] ||= []
    hash[list.val].push(list)
  end

  # puts "hash is #{hash}"
  while !hash.empty?
    k, min_list_arr = hash.min

    min_list_arr.each do |min_list|
      head.next = min_list
      head = head.next

      new_list = min_list.next
      if new_list
        hash[new_list.val] ||= []
        hash[new_list.val].push(new_list)
      end
    end

    hash.delete(k)

    # n_list = m_list.next
    # hash[n_list.val] = n_list if n_list
  end

  result.next
end

headers = []
[[1,4,5],[1,3,4],[2,6]].each do |arr|
  head = pre = ListNode.new(0)
  arr.each do |i|
    pre.next = ListNode.new(i)
    pre = pre.next
  end

  headers.push(head.next)
end

puts merge_k_lists(headers).inspect
