class MinStack

=begin
    initialize your data structure here.
=end
    def initialize()
      @array = []
      nil
    end


=begin
    :type x: Integer
    :rtype: Void
=end
    def push(x)
      min = @array.empty? ? x : [@array[-1][1], x].min
      @array.push([x, min])
      nil
    end


=begin
    :rtype: Void
=end
    def pop()
      v, _ = @array.pop
      v
    end


=begin
    :rtype: Integer
=end
    def top()
      @array[-1][0]
    end


=begin
    :rtype: Integer
=end
    def get_min()
      @array[-1][1]
    end


end

# Your MinStack object will be instantiated and called as such:
# obj = MinStack.new()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.get_min()

minStack = MinStack.new()
puts minStack.push(-2);
puts minStack.push(0);
puts minStack.push(-3);
puts minStack.get_min();   # --> 返回 -3.
puts minStack.pop();
puts minStack.top();      #--> 返回 0.
puts minStack.get_min();   #--> 返回 -2.
