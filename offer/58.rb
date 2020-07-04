class MaxQueue

  def initialize()
		@array = []
		@max_array = []
  end


=begin
  :rtype: Integer
=end
  def max_value()
		@max_array.empty? ? -1 : @max_array[-1]
  end


=begin
  :type value: Integer
  :rtype: Void
=end
  def push_back(value)
		@array.push(value)
		if @max_array.empty?
			@max_array.push(value)
		else
			while @max_array[-1] > value
				@max_array.pop
			end
			@max_array.push(value)
		end
  end


=begin
  :rtype: Integer
=end
  def pop_front()
		return -1 if result.empty?

		result = @array.shift

		if result == @max_array[0]
			@max_array.shift
		end

		result
  end
end

# Your MaxQueue object will be instantiated and called as such:
obj = MaxQueue.new()
param_1 = obj.max_value()
obj.push_back(value)
param_3 = obj.pop_front()