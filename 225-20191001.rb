class MyStack

=begin
    Initialize your data structure here.
=end
	def initialize()
		@primary_queue = Queue.new()
		@secondary_queue = Queue.new()
	end


=begin
    Push element x onto stack.
    :type x: Integer
    :rtype: Void
=end
	def push(x)
		@primary_queue.push(x)
	end


=begin
    Removes the element on top of the stack and returns that element.
    :rtype: Integer
=end
	def pop()
		(@primary_queue.size-1).times do
			@secondary_queue.push(@primary_queue.pop)
		end
		result = @primary_queue.pop

		@secondary_queue.size.times do
			@primary_queue.push(@secondary_queue.pop)
		end

		result
	end


=begin
    Get the top element.
    :rtype: Integer
=end
	def top()
		(@primary_queue.size-1).times do
			@secondary_queue.push(@primary_queue.pop)
		end
		result = @primary_queue.pop

		@secondary_queue.size.times do
			@primary_queue.push(@secondary_queue.pop)
		end
		@primary_queue.push(result)

		result
	end


=begin
    Returns whether the stack is empty.
    :rtype: Boolean
=end
    def empty()
			@primary_queue.size.zero?
    end


end

# Your MyStack object will be instantiated and called as such:
# obj = MyStack.new()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()