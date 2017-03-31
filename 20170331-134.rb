# @param {Integer[]} gas
# @param {Integer[]} cost
# @return {Integer}
#
# 如果中途失败了
# 则记录下还缺多少油
# 然后继续向前，记录下接下来的gas station能剩下多少gas
# 如果多余的gas 比之前缺少的要多，则返回 当前的result
def can_complete_circuit(gas, cost)
  gas_size = gas.size

  tank = 0
  total = 0
  result = 0
  gas_size.times do |start_pos|
    tank += (gas[start_pos] - cost[start_pos])

    if tank < 0
      total += tank
      tank = 0
      result = start_pos + 1
    end
  end

  (tank + total) < 0 ? -1 : result
end

