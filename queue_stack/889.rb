class MyStack

=begin
    Initialize your data structure here.
=end
    def initialize()
      @arr1 = []
      @arr2 = []
    end


=begin
    Push element x onto stack.
    :type x: Integer
    :rtype: Void
=end
    def push(x)
      @arr1.push(x)
    end


=begin
    Removes the element on top of the stack and returns that element.
    :rtype: Integer
=end
    def pop()
      (@arr1.size-1).times { @arr2.push(@arr1.shift) }
      result = @arr1.shift
      @arr1, @arr2 = @arr2, @arr1
      result
    end


=begin
    Get the top element.
    :rtype: Integer
=end
    def top()
      (@arr1.size-1).times { @arr2.push(@arr1.shift) }
      result = @arr2.push(@arr1.shift)
      @arr1, @arr2 = @arr2, @arr1
      result
    end


=begin
    Returns whether the stack is empty.
    :rtype: Boolean
=end
    def empty()
      @arr1.empty?
    end


end

# Your MyStack object will be instantiated and called as such:
# obj = MyStack.new()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()