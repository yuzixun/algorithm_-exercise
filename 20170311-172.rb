# Given an integer n, return the number of trailing zeroes in n!.
# Note: Your solution should be in logarithmic time complexity.

# 阶乘中要出现以10结尾的数字
# 必须是2*5出现
# 将数字分解后，2出现的次数必然大于5出现的次数
# 所以只需要求出 参与阶乘的数字中 5的个数
# 大于等于25的数，可能包含有多个5，所以采用递归，获取5的个数

# author: yuzixun
# @param {Integer} n
# @return {Integer}
def trailing_zeroes(n)
  n == 0 ? 0 : (n/5) + trailing_zeroes(n/5)
end