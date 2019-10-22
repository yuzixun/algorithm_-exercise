class LRUCache

=begin
    :type capacity: Integer
=end
    def initialize(capacity)
      @capacity = capacity
      @hash = {}
    end


=begin
    :type key: Integer
    :rtype: Integer
=end
    def get(key)
      value = @hash.delete(key)
      @hash[key] = value if value
      # puts "get key is #{key}, value is #{value}"
      value || -1
    end


=begin
    :type key: Integer
    :type value: Integer
    :rtype: Void
=end
    def put(key, value)
      @hash.delete(key)
      @hash[key] = value
      if @hash.size > @capacity
        first_key = @hash.first.first
        @hash.delete(first_key)
      end
      # puts "after put #{key}, #{value}, hash is #{@hash}"
      nil
    end


end

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache.new(capacity)
# param_1 = obj.get(key)
# obj.put(key, value)

