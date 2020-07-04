class Node
  attr_accessor :key, :value, :cnt

  def initialize(key, value, cnt)
    @key = key
    @value = value
    @cnt = cnt
  end
end

class LFUCache

=begin
    :type capacity: Integer
=end
    def initialize(capacity)
      @capacity = capacity
      @min_cnt = min_cnt
      @cnt_hash = {}
      @value_hash = {}
    end


=begin
    :type key: Integer
    :rtype: Integer
=end
    def get(key)
      return -1 if value_hash[key].nil?

      node = value_hash[key]
      @cnt_hash[node.cnt].delete(key)

      if @cnt_hash[node.cnt].empty?
        @cnt_hash.delete(node.cnt)
      end

      if @cnt_hash[@min_cnt].empty?
        @min_cnt+=1
      end

      node.cnt += 1
      @cnt_hash[node.cnt][key] ||= []
      @cnt_hash[node.cnt][key].push(node)
      return node.value
    end


=begin
    :type key: Integer
    :type value: Integer
    :rtype: Void
=end
    def put(key, value)
      return if @capacity <= 0
      if !@value_hash[key].nil?
        @value_hash[key].value = value
        return get(key)
      end

      if @value_hash.size >= @capacity
        key, node = @cnt_hash[@min_cnt].pop
        @value_hash.delete(key)
      end

      @value_hash[key] = @cnt_hash[1][key] = Node,new(key, value, 1)
      @min_cnt = 1
    end


end

# Your LFUCache object will be instantiated and called as such:
# obj = LFUCache.new(capacity)
# param_1 = obj.get(key)
# obj.put(key, value)