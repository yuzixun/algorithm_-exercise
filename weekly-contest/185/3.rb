# @param {String} croak_of_frogs
# @return {Integer}
def min_number_of_frogs(croak_of_frogs)
  return -1 if croak_of_frogs.size%5 != 0
  cache = {}
  croak_of_frogs.each_char do |c|
    cache[c] ||= 0
    cache[c]+=1
  end

  return -1 if cache.values.uniq.size != 1

  result = 0
  cache = Hash.new(0)
  croak_of_frogs.each_char do |c|
    case c
    when 'c'
      cache['c'] += 1
    when 'r'
      cache['r'] += 1
      return -1 if cache['c'] < cache['r']
    when 'o'
      cache['o'] += 1
      return -1 if cache['r'] < cache['o']
    when 'a'
      cache['a'] += 1
      return -1 if cache['o'] < cache['a']
    when 'k'
      cache['k'] += 1
      return -1 if cache['a'] < cache['k']
      "croak".each_char do |cc|
        cache[cc]-=1
      end
    end
    result = [result, cache['c']].max
  end
  result
end


p min_number_of_frogs("croakcroak")
p min_number_of_frogs("crcoakroak")
p min_number_of_frogs("croakcrook")
p min_number_of_frogs("croakcroa")

