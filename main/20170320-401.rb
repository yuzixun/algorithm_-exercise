# @param {Integer} num
# @return {String[]}
def read_binary_watch(num)
  binary =
    1023.times.map do |i|
      data = "%010b" % i
      data if data.count('1') == num
    end.compact

  binary.reduce([]) do |results, data|
    hour, minute = data[0..3].to_i(2), data[4..-1].to_i(2)
    (hour > 11 or minute > 59) ? results : results.push("#{hour}:#{'%02d' % minute}")
  end
end
