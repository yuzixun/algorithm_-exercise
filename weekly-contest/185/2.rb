# @param {String[][]} orders
# @return {String[][]}
def display_table(orders)
  tables = {}
  cais = []

  orders.each do |order|
    cais.push(order[2])

    tables[order[1].to_i] ||= {}
    tables[order[1].to_i][order[2]] ||= 0
    tables[order[1].to_i][order[2]] += 1
  end

  result = []
  cais.uniq!
  cais.sort!
  result[0] = ["Table"] + cais

  tables.sort.each_with_index do |table_with_os, index|
    table, cais_with_count = table_with_os
    cur = [table.to_s]
    cais.each_with_index do |cai, index|
      cur[index+1] = cais_with_count[cai].to_i.to_s
    end

    result[index+1] = cur
  end

  result
end



orders = [["David","3","Ceviche"],["Corina","10","Beef Burrito"],["David","3","Fried Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
p display_table(orders)
orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian Waffles"],["Brianna","1","Canadian Waffles"]]
p display_table(orders)
orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef Burrito"],["Melissa","2","Soda"]]
p display_table(orders)
