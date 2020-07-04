class SortedStack
    def initialize()
        @stack1 = []
        @stack2 = []
    end


=begin
    :type val: Integer
    :rtype: Void
=end
    def push(val)
      while @stack1[-1] >= val
        @stack2.push(@stack1.pop)
      end
      @stack1.push(val)
      while !@stack2.empty?
        @stack1.push(@stack2.pop)
      end
    end


=begin
    :rtype: Void
=end
    def pop()
      @stack1.pop
    end


=begin
    :rtype: Integer
=end
    def peek()
      @stack1[-1]
    end


=begin
    :rtype: Boolean
=end
    def is_empty()
      @stack1.empty?
    end


end

# Your SortedStack object will be instantiated and called as such:
# obj = SortedStack.new()
# obj.push(val)
# obj.pop()
# param_3 = obj.peek()
# param_4 = obj.is_empty()