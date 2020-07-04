def find_half

end


def insert_into(sorted_array, insert_element)
  will_replace_at = nil
  will_insert_at = nil

  sorted_array

  sorted_array.each_with_index do |element, index|
    if element.nil?
      will_insert_at = index
      break
    end

    if insert_element > element
      will_insert_at = index + 1
    else
      break
    end
  end
end

def get_kth_largest_element_by(array, k)
  return if array.size < k

  sorted_array = Array.new(k)

  array.each do |element|
    sorted_array = insert_into(sorted_array, element)
    sorted_array.shift if sorted_array.size > k
  end

  puts "sorted_array si #{sorted_array}"
  array.first
end


arr = [3, 2, 1, 5, 6, 4]
k = 2

get_kth_largest_element_by(arr, k)