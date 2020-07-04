require 'date'
# @param {String} date1
# @param {String} date2
# @return {Integer}
def days_between_dates(date1, date2)
  (Date.parse(date1) - Date.parse(date2)).to_i.abs
end

date1 = "2019-06-29"; date2 = "2019-06-30"
puts days_between_dates(date1, date2)
date1 = "2020-01-15"; date2 = "2019-12-31"
puts days_between_dates(date1, date2)

