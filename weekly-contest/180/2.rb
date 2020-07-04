class CustomStack

=begin
    :type max_size: Integer
=end
    def initialize(max_size)
      @arr = []
      @max = max_size
    end


=begin
    :type x: Integer
    :rtype: Void
=end
    def push(x)
      return if @arr.size >= @max
      @arr.push(x)
    end


=begin
    :rtype: Integer
=end
    def pop()
      @arr.empty? ? -1 : @arr.pop
    end


=begin
    :type k: Integer
    :type val: Integer
    :rtype: Void
=end
    def increment(k, val)
      size  = @arr.size
      @arr = @arr.map.with_index { |i, index| index < k ? i+val : i}
    end


end

puts customStack = CustomStack.new(3); # 栈是空的 []
puts customStack.push(1);                          # 栈变为 [1]
puts customStack.push(2);                          # 栈变为 [1, 2]
puts customStack.pop();                            # 返回 2 --> 返回栈顶值 2，栈变为 [1]
puts customStack.push(2);                          # 栈变为 [1, 2]
puts customStack.push(3);                          # 栈变为 [1, 2, 3]
puts customStack.push(4);                          # 栈仍然是 [1, 2, 3]，不能添加其他元素使栈大小变为 4
puts "customStack.increment(5, 100); is #{customStack.increment(5, 100);}"                # 栈变为 [101, 102, 103]
puts "customStack.increment(2, 100); is #{customStack.increment(2, 100);}"                # 栈变为 [201, 202, 103]
puts customStack.pop();                            # 返回 103 --> 返回栈顶值 103，栈变为 [201, 202]
puts customStack.pop();                            # 返回 202 --> 返回栈顶值 202，栈变为 [201]
puts customStack.pop();                            # 返回 201 --> 返回栈顶值 201，栈变为 []
puts customStack.pop();                            # 返回 -1 --> 栈为空，返回 -1