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
      value = @hash[key]
      if value
        @hash.delete(key)
        @hash[key] = value
      end
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
      hash.shift if hash.size > @capacity

      nil
    end


end

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache.new(capacity)
# param_1 = obj.get(key)
# obj.put(key, value)