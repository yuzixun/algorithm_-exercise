# @param {String} str
# @return {Integer}
def str_to_int(str)
  [[str.to_i, limit = -2**31].max, -limit-1].min
end

