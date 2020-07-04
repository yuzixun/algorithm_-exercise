class MyQueue

=begin
    Initialize your data structure here.
=end
    def initialize()
      @stack1 = []
      @stack2 = []
    end


=begin
    Push element x to the back of queue.
    :type x: Integer
    :rtype: Void
=end
    def push(x)
      @stack1.push(x)
    end


=begin
    Removes the element from in front of queue and returns that element.
    :rtype: Integer
=end
    def pop()
      (@stack1.size-1).times { @stack2.push(@stack1.pop) }
      result = @stack1.pop
      @stack2.size.times { @stack1.push(@stack2.pop) }
      result
    end


=begin
    Get the front element.
    :rtype: Integer
=end
    def peek()
      (@stack1.size-1).times { @stack2.push(@stack1.pop) }
      result = @stack1.pop
      @stack2.size.times { @stack1.push(@stack2.pop) }
      @stack1.push(result)
      result
    end


=begin
    Returns whether the queue is empty.
    :rtype: Boolean
=end
    def empty()
      @stack1.empty?
    end


end

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue.new()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()