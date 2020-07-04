class MyQueue

=begin
    Initialize your data structure here.
=end
    def initialize()
      @in = []
      @out = []
    end


=begin
    Push element x to the back of queue.
    :type x: Integer
    :rtype: Void
=end
    def push(x)
      @in.push(x)
    end


=begin
    Removes the element from in front of queue and returns that element.
    :rtype: Integer
=end
    def pop()
      puts "poping #{@in}, #{@out}"
      (@in.size-1).times do
        @out.push(@in.pop)
      end
      puts "@in is #{@in}, @out is #{@out}"
      result = @in.pop

      @out.size.times do
        @in.push(@out.pop)
      end

      result
    end


=begin
    Get the front element.
    :rtype: Integer
=end
    def peek()
      (@in.size-1).times do
        @out.push(@in.pop)
      end

      result = @in.pop

      @in.push(result)
      (@out.size).times do
        @in.push(@out.pop)
      end

      result
    end


=begin
    Returns whether the queue is empty.
    :rtype: Boolean
=end
    def empty()
      @in.empty?
    end


end

# Your MyQueue object will be instantiated and called as such:
# queue = MyQueue.new()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
