# author: yuzixun
# @param {String} num
# @param {Integer} k
# @return {String}

def remove_kdigits(num, k)
  return '0' if num.size == k

  result = ''
  num.each_char do |element|
    while k > 0 && result.slice(-1) && result.slice(-1) > element
      result.chop!
      k -= 1
    end
    result.concat(element)
  end

  result = result[0..result.size-1-k].sub(/^0*/, '')
  result.empty? ? '0' : result
end
